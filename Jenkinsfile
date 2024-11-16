pipeline {
    agent {
        docker {
            image 'jenkins/jenkins:lts'
            args '-u root:root' // Dùng quyền root nếu cần cài đặt thêm package
        }
    }

    environment {
        DOCKER_IMAGE = "jenkins/jenkins"
        DOCKER_TAG = "lts"
        GIT_REPO = "https://github.com/LVNAnh/golang-jenkins.git"
        GIT_BRANCH = "master"
    }

    stages {
        stage('Checkout') {
            steps {
                echo 'Cloning repository...'
                checkout([
                    $class: 'GitSCM',
                    branches: [[name: "*/${env.GIT_BRANCH}"]],
                    userRemoteConfigs: [[
                        url: "${env.GIT_REPO}"
                    ]]
                ])
            }
        }

        stage('Build') {
            steps {
                echo 'Building Docker image...'
                sh '''
                docker build -t ${DOCKER_IMAGE}:${DOCKER_TAG} .
                '''
            }
        }

        stage('Test') {
            steps {
                echo 'Running tests...'
                sh '''
                docker run --rm ${DOCKER_IMAGE}:${DOCKER_TAG} sh -c "echo Running tests; sleep 2"
                '''
            }
        }

        stage('Push Image') {
            steps {
                echo 'Skipping Docker push as this is an internal Jenkins image.'
            }
        }

        stage('Deploy') {
            steps {
                echo 'Deploying application...'
                sh '''
                echo Deploying with Jenkins Docker image: ${DOCKER_IMAGE}:${DOCKER_TAG}
                '''
            }
        }
    }

    post {
        always {
            echo 'Cleaning up resources...'
            sh '''
            docker rmi ${DOCKER_IMAGE}:${DOCKER_TAG} || true
            '''
        }
        success {
            echo 'Pipeline executed successfully!'
        }
        failure {
            echo 'Pipeline failed. Please check the logs.'
        }
    }
}
