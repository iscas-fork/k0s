# 定义URL数组
urls=(
  "api"
  "apiextensions-apiserver"
  "apimachinery"
  "cli-runtime"
  "client-go"
  "cloud-provider"
  "component-base"
  "component-helpers"
  "cri-api"
  "kube-aggregator"
  "kubectl"
  "kubelet"
  "kubernetes"
  "mount-utils"
  "utils"
)

# 遍历数组
for url in "${urls[@]}"; do
    curl="https://raw.githubusercontent.com/kubernetes/"$url"/main/go.mod"
    durl="https://raw.githubusercontent.com/kubernetes/"$url"/master/go.mod"
    echo "k8s.io/"$url
    # 使用curl获取URL内容，并打印出来
    # echo $curl"\n"
    curl -s $curl
    echo ""
    # echo $durl"\n"
    curl -s $durl
    echo "----------------------------------"
done
