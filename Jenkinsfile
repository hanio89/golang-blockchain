pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                // Checkout source code from version control (e.g., Git)
                checkout scm
            }
        }

        stage('Build') {
            steps {
                // Build the Golang application
                script {
                    // Build the Golang application
                    sh 'go run main.go'
                }
            }
        }

    }

    post {
        success {
            // This block is executed if the pipeline is successful
            echo 'Pipeline completed successfully!'
        }

        failure {
            // This block is executed if the pipeline fails
            echo 'Pipeline failed!'
        }
    }
}
