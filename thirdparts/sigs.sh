# 定义URL数组
urls=(
  "controller-runtime"
  "yaml"
)

# 遍历数组
for url in "${urls[@]}"; do
    curl="https://raw.githubusercontent.com/kubernetesi-sigs/"$url"/main/go.mod"
    durl="https://raw.githubusercontent.com/kubernetes-sigs/"$url"/master/go.mod"
    echo "k8s.io/"$url
    # 使用curl获取URL内容，并打印出来
    # echo $curl"\n"
    curl -s $curl
    echo ""
    # echo $durl"\n"
    curl -s $durl
    echo "----------------------------------"
done
