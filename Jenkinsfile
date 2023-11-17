pipeline {
    agent any

    stages {
        stage('Debug') {
            steps {
                script {
                    echo "Jenkins Home: ${JENKINS_HOME}"
                    echo "Jenkins Workspace: ${WORKSPACE}"
                    echo "Java Home: ${JAVA_HOME}"
                    echo "PATH: ${PATH}"
                    echo "Current User: ${USER}"
                }
            }
        }

        // Add other stages
    }

    // Add post-build actions if necessary
}