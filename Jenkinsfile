pipeline {
    agent any

    environment {
        GO_IMAGE_NAME = "getgoal-app"
        CONTAINER_NAME = "getgoal-app"
    }

    stages {
      stages {
        stage('Checkout') {
            steps {
                script {
                    checkout([$class: 'GitSCM', branches: [[name: 'refs/heads/*']], userRemoteConfigs: [[url: 'https://github.com/GetGoal/GetGoal-app.git']]])
                }
            }
        }
      }

        stage('Build App Images | Check ENV') {
            steps {
                script {
                    sh "echo ${ENV}"
                    sh "echo ${GIT_TAG}"
                    // sh "docker build -t ${POSTGRES_IMAGE_NAME}:${IMAGE_TAG} ."
                }
            }
        }
        // stage ('Remove container'){
        //     steps {
        //       script {
        //             // Run the command and capture the exit code
        //             def exitCode = sh(script: "docker rm -f ${CONTAINER_NAME}-${params.deployEnvironment}", returnStatus: true)

        //             // Check the exit code to determine success or failure
        //             if (exitCode == 0) {
        //                 echo "Container removal was successful"
        //                 // Add more steps or logic here if needed
        //             } else {
        //                 echo "Container removal failed or was skipped"
        //                 // Add more steps or logic here if needed
        //             }
        //       }
        //     }
        // }

        // stage('Deploy') {
        //     steps {
        //         script {
        //           sh "docker run -d -v ${HOST_PATH}${params.deployEnvironment}:/var/lib/postgresql/data  --name ${CONTAINER_NAME}-${params.deployEnvironment} ${POSTGRES_IMAGE_NAME}:${IMAGE_TAG}"
        //         }
        //     }
        // }

        // stage('Link Networks') {
        //     steps {
        //         script {

        //           sh "docker network connect ${params.deployEnvironment}-network ${CONTAINER_NAME}-${params.deployEnvironment}"
        //         }
        //     }
        
        // }

        // stage('Clear Storage') {
        //     steps {
        //         script {
        //             sh "docker image prune -a -f"
        //         }
        //     }
        // }

        // stage('Health Cheack') {
        //     steps {
        //         script {
        //             def containerId = sh(script: "docker ps -q --filter name=${CONTAINER_NAME}-${params.deployEnvironment}", returnStdout: true)

        //             if (containerId) {
        //                 def healthStatus = sh(script: "docker inspect --format '{{.State.Running}}'  ${containerId}", returnStdout: true)
                        
        //                 echo "Helath : ${healthStatus}"
        //                 if (healthStatus) {
        //                     echo "Container is running healthily."
        //                 } else {
        //                     error "Unable to retrieve container health status."
        //                 }
        //             } else {
        //                 error "Container not found. Make sure it is running."
        //             }
        //         }
        //     }
        // }
    }

    post {
        success {
            echo 'Pipeline successfully completed!'
        }
        failure {
            echo 'Pipeline failed!'
        }
    }
}