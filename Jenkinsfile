pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                // Checkout source code from version control (e.g., Git)
                checkout scm
            }
        }

        stage('Test') {
            steps {
                // Run tests using go test
                script {
                    sh 'go test ./...'
                }
            }
        }

        stage('Build') {
            steps {
                // Build the Golang application
                script {
                    sh 'go build -o myapp'
                }
            }
        }

        stage('Deploy') {
            steps {
                // Deploy the application (you can customize this based on your deployment process)
                script {
                    sh './myapp'
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
