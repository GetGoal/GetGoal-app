pipeline {
    agent any

    stages {
        stage('Debug') {
            steps {
                script {
                    echo "Jenkins Home: ${JENKINS_HOME}"
                    echo "Jenkins Workspace: ${WORKSPACE}"
                    echo "GIT_TAG: ${GIT_TAG}"
                    echo "Environtment to depliy: ${ENV}"
            }
        }

        // Add other stages
    }

    // Add post-build actions if necessary
}