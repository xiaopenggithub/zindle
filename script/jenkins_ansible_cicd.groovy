#!groovy

pipeline{
    agent {node {label 'node98'}}

    environment{
        PATH="/bin:/sbin:/usr/bin:/usr/sbin:/usr/local/bin"
    }
    
    parameters{
        choice(
            choices: 'test\nprod',
            description: 'choose deploy environment',
            name: 'deploy_env'
        )        
        string(name: 'version', defaultValue: 'v1.0.0', description: 'build version') 
        string(name: 'branch', defaultValue: 'master', description: 'repo branch') 
        booleanParam(
            name: 'is_no_cache',
            defaultValue: false,
            description: 'build no with no cache'
        )
        booleanParam(
            name: 'publish_after_image_builded',
            defaultValue: true,
            description: 'Deploy image to server'
        )
    }

    stages {
        stage('Checkout Backend repo') {
            steps{
                sh 'git config --global http.sslVerify false'
                dir ("${env.WORKSPACE}") {
                    git branch: "$branch",credentialsId: "f5791b95-6e08-4e67-b842-43843b9d7282", url:'http://127.0.0.1/qinrongqiang/repo.git'
                }
            }
        }

        stage("Build Backend Image"){
            steps {
                dir ("${env.WORKSPACE}") {
                    sh """
                    if ${is_no_cache}
                    then                        
                        docker build --no-cache -t harbor.qin.com/repo/psmis_backend:$version -f ./dockerfiles/Dockerfile$deploy_env .
                    else
                        docker build -t harbor.qin.com/repo/psmis_backend:$version -f ./dockerfiles/Dockerfile$deploy_env .
                    fi
                    echo "[INFO] Build Backend Image Done..." 
                    """
                }
            }
        }

        stage("Push Backend Image to Harbor"){
            steps {
            	dir ("${env.WORKSPACE}") {
            		sh """
                    docker push harbor.qin.com/repo/psmis_backend:$version
            		echo "[INFO] Push Backend Image to Harbor Done..." 
                    """
            	}
            }
        }
        
        stage("Deploy Backend Image"){
            steps {
                sh"""
				set +x
                if ${publish_after_image_builded}
                then
              		echo "[INFO] Deploying Backend Image"
              		cd /home/user/py3
                    source ./bin/activate
                    source ./ansible/hacking/env-setup -q
                    
                    ansible --version
                    ansible-playbook --version
                    
                    cd /home/user/py3/psmis_playbooks
                    ansible-playbook -i inventory/testenv ./deploy.yml
                    echo "[INFO] Deploy Backend Image Done..."
                else
              		echo "[INFO] Only build Backend,Done..."
                fi
                set -x
                """
            }
        }
    }
}