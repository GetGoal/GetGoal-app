pipeline {
    agent any
      environment {
        IMAGE_NAME = "getgoal-app"
        CONTAINER_NAME = "getgoal-app"
    }
    stages {

        stage('Build GO-App Images') {
            steps {
                script {
                    sh "docker build \
                    --build-arg ENV=${ENV} \
                    --build-arg DB_HOST=${env.DB_HOST}-${ENV} \
                    --build-arg DB_USER=${env.DB_USER}  \
                    --build-arg DB_PORT=${env.DB_PORT}  \
                    --build-arg DB_NAME=${env.DB_NAME}  \
                    --build-arg DB_PASSWORD=${DB_PASSWORD} \
                    -t ${IMAGE_NAME}:${GIT_TAG} ."
                }
            }
        }
    
        stage ('Remove container'){
            steps {
              script {
                    // Run the command and capture the exit code
                    def exitCode = sh(script: "docker rm -f ${CONTAINER_NAME}-${ENV}", returnStatus: true)

                    // Check the exit code to determine success or failure
                    if (exitCode == 0) {
                        echo "Container removal was successful"
                        // Add more steps or logic here if needed
                    } else {
                        echo "Container removal failed or was skipped"
                        // Add more steps or logic here if needed
                    }
              }
            }
        }

        stage('Deploy and link network') {
            steps {
                script {
                  sh "docker run -d \
                  --name ${CONTAINER_NAME}-${ENV} \
                  --hostname ${CONTAINER_NAME}-${ENV} \
                  --network  ${ENV}-network \
                  --restart unless-stopped ${IMAGE_NAME}:${GIT_TAG}"
                }
            }
        }

        stage('Clear Storage') {
            steps {
                script {
                    sh "docker image prune -a -f"
                }
            }
        }

        stage('Health Cheack') {
            steps {
                script {
                    def containerId = sh(script: "docker ps -q --filter name=${CONTAINER_NAME}-${ENV}", returnStdout: true)

                    if (containerId) {
                        def healthStatus = sh(script: "docker inspect --format '{{.State.Running}}'  ${containerId}", returnStdout: true)
                        
                        echo "Helath : ${healthStatus}"
                        if (healthStatus) {
                            echo "Container is running healthily."
                        } else {
                            error "Unable to retrieve container health status."
                        }
                    } else {
                        error "Container not found. Make sure it is running."
                    }
                }
            }
        }
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