#!groovy

void setBuildStatus(String message, String state) {
  step([
      $class: "GitHubCommitStatusSetter",
      reposSource: [$class: "ManuallyEnteredRepositorySource", url: "https://github.com/kibbles-n-bytes/service-catalog"],
      contextSource: [$class: "ManuallyEnteredCommitContextSource", context: "kibbles-n-bytes-testing"],
      errorHandlers: [[$class: "ChangingBuildStatusErrorHandler", result: "UNSTABLE"]],
      statusResultSource: [ $class: "ConditionalStatusResultSource", results: [[$class: "AnyBuildResult", message: message, state: state]] ]
  ]);
}

node {
  echo "Hello world1."
  setBuildStatus("I'm done", "SUCCESS");
}
