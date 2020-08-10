/*
Author: zuoguocai@126.com
Tips: 1. 注意单引号和双引号的区别
      2. git tag ---docker tag --- k8s deployment image tag  做的关联
      3. 少了单元测试，代码覆盖，质量管理，质量门禁，度量等
      4. 少了异常捕获，少了多分支环境切换
      5. harbor 和 k8s worker node 上需要回收不用的image，docker image prune -a --force，最好加上label，避免误删除

*/
pipeline {
    agent any
    
    tools {
        //maven 'mvn3.6'
        //jdk   'java1.8'
        go  'go1.13'
    }
    
    environment {
    
        HARBOR_CREDS = credentials('jenkins-harbor-creds')
        BUILD_USER_ID = ""
        BUILD_USER = ""
        BUILD_USER_EMAIL = ""
        
    }
    
    stages {
    
    
        stage('准备环境变量'){
        
              steps {
              // 由插件user build vars 提供
              
               wrap([$class: 'BuildUser']) {
                   script {
                       BUILD_USER_ID = "${env.BUILD_USER_ID}"
                       BUILD_USER = "${env.BUILD_USER}"
                       BUILD_USER_EMAIL = "${env.BUILD_USER_EMAIL}"
                   }
				}
				// Test out of wrap
				echo "Build User ID: ${BUILD_USER_ID}"
				echo "Build User: ${BUILD_USER}"
				echo "Build User Email: ${BUILD_USER_EMAIL}"
            }
        
        }    
    
    
        stage('拉取代码') { // for display purposes
        
            steps{
                 // 清理工作区
                 step([$class: 'WsCleanup'])
                 // 拉取代码
                //checkout([$class: 'GitSCM', branches: [[name: '*/master']], doGenerateSubmoduleConfigurations: false, extensions: [], submoduleCfg: [], userRemoteConfigs: [[credentialsId: '3519b4b1-47fd-43e7-bec3-502aa5d5a99a', url: 'https://git.zuoguocai.io/openstack/getrealip.git']]])
            
                //GIT_TAG = sh(returnStdout: true,script: 'git describe --tags --always').trim()
                
                
                  
                 //git branch: '${BUILD_BRANCH}',
                
                 script {
                 
                 git credentialsId: '3519b4b1-47fd-43e7-bec3-502aa5d5a99a', url: "https://git.zuoguocai.io/openstack/getrealip.git"
                 
                
                  build_tag = sh(returnStdout: true, script: 'git describe --tags --always').trim()
                 //echo build_tag
                 
                 }
                
                
            }
        }
        
        stage('编译代码') {
            steps {
                sh 'go version'
                sh 'go build .'
            }
        }
        
 
         
        
        
        stage('构建镜像'){
         
            input {
                message "测试环境"
                ok "提交."
                submitter ""
                parameters {
                    string(name: 'PASSWD', defaultValue: '', description: '请输入密码开始部署')
                }
            }
         
            steps {
            
              script{
              if (PASSWD == HARBOR_CREDS_PSW) {
                echo "start build image"
                
                  dir('') {
                    // 删除之前构建镜像
                    sh "docker image prune -a --force  --filter 'label=ZuoGuocai'"
                    // build镜像
                    
                    //echo build_tag
                    
                   sh "docker build -t registry.zuoguocai.io/release/getrealip:${build_tag} ."
                    // 登录镜像仓库
                    sh "docker login -u ${HARBOR_CREDS_USR} -p ${HARBOR_CREDS_PSW} registry.zuoguocai.io"
                    // 推送镜像到镜像仓库
                   sh "docker push registry.zuoguocai.io/release/getrealip:${build_tag}"
                   
                   }
                
                
                 } else {
                     echo '密码错误,部署失败'
                  }
                }
            }
          
        }
        
            
        stage('部署到k8s集群') {
              
               steps{
                    withKubeConfig([credentialsId: 'k8s-token', serverUrl: 'https://k8s-api.zuoguocai.io:6443']) {
                    
                    echo build_tag
                  
                    sh "sed  's/<IMG_TAG>/${build_tag}/g' /opt/deploy/getrealip.temp   > /opt/deploy/getrealip.yaml"
                    sh "kubectl apply  -f /opt/deploy/getrealip.yaml"
                    sh "kubectl get pods -n devops"
                                       
                    }
                    
                }
                
                post {
                 success{
                    dingtalk (
                    // robot 为插件DingTalk配置后自动生成的id,在系统管理--系统配置--钉钉里找
                        robot: '14205481-9c8d-40d3-a667-95a91a09b33f',
                        type: 'MARKDOWN',
                        title: 'Jenkins pipeline构建通知',
                        text: [
                            "# <font color=#66CDAA>${env.BUILD_DISPLAY_NAME}构建成功 </font>",
                           '---',
                           "- 执行人: ${BUILD_USER}",
                           "- 邮箱: ${BUILD_USER_EMAIL}",
                           "- 作业: ${env.WORKSPACE}",
                          
                        ],
                        at: [
                          '130xxxxxx'
                        ]
                    )
                }
            }
        }
        
       
        

    }
}
