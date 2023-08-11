node {
    def root = tool type: 'go', name: '1.20'

    withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
        sh 'go version'
    }
}