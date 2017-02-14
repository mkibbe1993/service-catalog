#!groovy

/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Jenkins continuous integration
//
// Parameters Jenkins needs to / can supply:
//
// TEST_PROJECT:   Google Cloud Project ID of the project to use for testing.
// TEST_ZONE:      GCP Zone in which to create test GKE cluster
// TEST_ACCOUNT:   GCP service account credentials (JSON file) to use for testing.

def namespace    = 'catalog'
def root_path    = 'src/github.com/kubernetes-incubator/service-catalog'

node {
  echo "Service Catalog end-to-end test"

  // Checkout the source code.
  checkout scm

  env.GOPATH = env.WORKSPACE
  env.ROOT = "${env.WORKSPACE}/${root_path}"
  env.KUBECONFIG = "${env.ROOT}/kubeconfig"

  dir([path: env.ROOT]) {
    // Run build.

    def clustername = "jenkins-" + sh([returnStdout: true, script: '''openssl rand \
        -base64 100 | tr -dc a-z0-9 | cut -c -25''']).trim()

    try {
      // Initialize build, for example, updating installed software.
      sh """${env.ROOT}/contrib/jenkins/init_build.sh"""

      // These are done in parallel since creating the cluster takes a while,
      // and the build doesn't depend on it.
      sh """${env.ROOT}/contrib/jenkins/build.sh --no-docker-compile \
	  --project ${test_project} \
	  --coverage '${env.WORKSPACE}/coverage.html'"""
    } catch (Exception e) {
      currentBuild.result = 'FAILURE'
    }

    if (currentBuild.result == 'FAILURE') {
      error 'Build failed.'
    }
  }

  archiveArtifacts artifacts: 'coverage.html', allowEmptyArchive: true, onlyIfSuccessful: true
}
