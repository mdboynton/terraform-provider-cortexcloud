Tools used:
  - Go v1.24.3
  - Terraform v1.11.4
  - Copywrite v0.22.0
    - `brew tap hashicorp/tap`
    - `brew install hashicorp/tap/copywrite`
  - Tfplugindocs v0.21.0
    - `go get github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs`


## Bash Aliases

Below are some bash aliases to assist with testing and debugging the provider.

* Paste the code below into your `.bashrc` file (or analogous configuration file for your terminal of choice).
* Update the `providerDirPath` value in the `buildProvider()` function with the full path to the root provider directory on your local machine
* Apply your updated terminal configuration by running `source $HOME/.bashrc` (replace with the path to your terminal's configuration file)

This will create the following aliases:
* `tf-build` - Build the provider and replace the binary in your Terraform plugins directory with the newly built binary file
* `tf-apply` - Build the provider, remove all Terraform directories (including state file) and run a `terraform apply` operation with debug logs enabled
* `tf-apply-state` - Same as `tf-apply` except the state file is preserved
  * Use this when testing provider code that is intended to work with existing resources or data sources in your Terraform state file (e.g. updating a resource)
* `tf-apply-nb` - Same as `tf-apply` except the provider binary is not re-built
* `tf-plan` - Build the provider, remove all Terraform directories (including state file) and run a `terraform plan` operation with debug logs enabled
* `tf-plan-state` - Same as `tf-plan` except the state file is preserved
* `tf-plan-nb` - Same as `tf-plan` except the provider binary is not re-built


```
setopt extended_glob # enables the use of ^ to negate pattern in rm

# Detect if debugger is running and set TF_REATTACH_PROVIDERS environment variable
function handleAttachTerraformToDebugger() {
    processes=$(ps -A)

    delveRunning=$(echo "${processes}" | grep "dlv")
    if [ -z "${delveRunning}" ]; then
        echo "Delve not running"
        if ! [[ -z "${TF_REATTACH_PROVIDERS}" ]]; then
            echo "Clearing TF_REATTACH_PROVIDERS environment variable"
            export TF_REATTACH_PROVIDERS=''
        fi
        return
    fi

    echo "Debug session detected. Attaching Terraform CLI..."
    
    providerDirName="terraform-provider-cortexcloud"
    providerBinaryName=$(echo "${providerDirName}" | awk -F '-' '{print $NF}')
    echo "providerBinaryName = ${providerBinaryName}"

    debugPID="$(ps -A | grep "${providerDirName}/__debug" | grep -v /Library | grep -v grep | awk '{print $1}')"
    echo "debugPID = ${debugPID}"

    socket=$(lsof -p $debugPID | grep /var/folders/ | awk '{print $NF}')
    echo "socket = ${socket}"

    envVarValue="{\"registry.terraform.io/PaloAltoNetworks/${providerBinaryName}\":{\"Protocol\":\"grpc\",\"ProtocolVersion\":6,\"Pid\":${debugPID},\"Test\":true,\"Addr\":{\"Network\":\"unix\",\"String\":\"${socket}\"}}}"

    echo "Setting TF_REATTACH_PROVIDERS=\"${envVarValue}\""
    export TF_REATTACH_PROVIDERS="${envVarValue}"
}

function deleteTerraformState() {
    echo "Deleting terraform.tfstate"
    rm -f terraform.tfstate
}

function deleteTerraformModules() {
    echo "Deleting .terraform/modules/"
    rm -rf .terraform/modules
}

function deleteAllTerraformFiles() {
    rmdir -r .terraform*
    rmdir -r terraform*
}

function cleanTerraform() {
    echo "Deleting cortex cloud provider from .terraform.lock.hcl"
    sed -i '' '/^provider "registry.terraform.io\/paloaltonetworks\/cortexcloud" {/,/^}/d' .terraform.lock.hcl

    # Delete state file if first argument is 1 
    if [[ $1 == 1 ]]; then
        deleteTerraformState
    fi

    deleteTerraformModules
}

function terraformInitAndApply() {
    terraform init
    TF_LOG=DEBUG TMPDIR="$PWD" terraform apply -auto-approve 2>&1 | grep -v "Value switched" | grep -v "Marking computed attributes"
}

function terraformInitAndPlan() {
    terraform init
    TF_LOG=DEBUG TMPDIR="$PWD" terraform plan 2>&1 | grep -v "Value switched" | grep -v "Marking computed attributes"
}

function tfApply() {
    cleanResult=$(cleanTerraform 1)
    if [[ $cleanResult != 1 ]]; then
        terraformInitAndApply
    fi
}

function tfApplyPreserveState() {
    cleanResult=$(cleanTerraform 0)
    if [[ $cleanResult != 1 ]]; then
        terraformInitAndApply
    fi
}

function tfPlan() {
    cleanResult=$(cleanTerraform 1)
    if [[ $cleanResult != 1 ]]; then
        terraformInitAndPlan
    fi
}

function tfPlanPreserveState() {
    cleanResult=$(cleanTerraform 0)
    if [[ $cleanResult != 1 ]]; then
        terraformInitAndPlan
    fi
}

function buildProvider() {
    # Update this value with the path to the provider directory on your local machine
    providerDirPath="$HOME/terraform-provider-cortexcloud"
    
    ( cd "$providerDirPath" && make ) 
    buildStatus=$?

    if [[ $buildStatus -ne 0 ]]; then
        echo "Provider build failed."
        return 1
    else
        echo "Provider successfully built"
        return 0
    fi
}

function tfEntrypoint() {
    handleAttachTerraformToDebugger

    if [[ $2 == 1 ]]; then
        buildProvider
        local buildExitCode=$?

        if [[ $buildExitCode -ne 0 ]]; then
            return $?
        fi
    fi
    
    if [[ $1 == 1 ]]; then
        tfApply
    elif [[ $1 == 2 ]]; then
        tfApplyPreserveState
    elif [[ $1 == 3 ]]; then
        tfPlan
    elif [[ $1 == 4 ]]; then
        tfPlanPreserveState
    fi

    return 0
}

alias tf-build="buildProvider"
alias tf-apply="tfEntrypoint 1 1"
alias tf-apply-nb="tfEntrypoint 1 2"
alias tf-apply-state="tfEntrypoint 2 1"
alias tf-plan="tfEntrypoint 3 1"
alias tf-plan-nb="tfEntrypoint 3 2"
alias tf-plan-state="tfEntrypoint 4 1"
```


## Debugging with Visual Studio Code

To debug the provider, install the [Go for Visual Studio Code](https://marketplace.visualstudio.com/items?itemName=golang.Go) extension from the marketplace.

After installing the extension, open the project in Visual Studio Code select the `Run and Debug` option on the sidebar and click `create a launch.json file`.

Example `launch.json` definition:

```
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Debug Terraform Provider",
            "type": "go",
            "request": "launch",
            "mode": "debug",
            // this assumes your workspace is the root of the repo
            "program": "${workspaceFolder}",
            "env": {},
            "args": [
                "-debug=true",
                ">",
                "${workspaceFolder}/debug_output.txt",
            ],
        }
    ]
}
```
