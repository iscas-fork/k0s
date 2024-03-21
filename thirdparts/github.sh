# 定义URL数组
urls=(
  "BurntSushi/toml"
  "Masterminds/semver"
  "Masterminds/sprig"
  "asaskevich/govalidator"
  "avast/retry-go"
  "bombsimon/logrusr"
  "carlmjohnson/requests"
  "cavaliergopher/grab/v3"
  "cilium/ebpf"
  "cloudflare/cfssl"
  "containerd/cgroups"
  "containerd/containerd"
  "denisbrodbeck/machineid"
  "evanphx/json-patch/v5"
  "fsnotify/fsnotify"
  "go-logr/logr"
  "go-openapi/jsonpointer"
  "go-playground/validator"
  "google/go-cmp"
  "hashicorp/terraform-exec"
  "k0sproject/bootloose"
  "k0sproject/dig"
  "k0sproject/version"
  "kardianos/service"
  "logrusorgru/aurora"
  "mesosphere/toml-merge"
  "mitchellh/go-homedir"
  "olekukonko/tablewriter"
  "opencontainers/runtime-spec"
  "otiai10/copy"
  "pelletier/go-toml"
  "robfig/cron"
  "rqlite/rqlite"
  "segmentio/analytics-go"
  "sirupsen/logrus"
  "spf13/cobra"
  "spf13/pflag"
  "stretchr/testify"
  "urfave/cli"
  "vishvananda/netlink"
  "vmware-tanzu/sonobuoy"
  "zcalusic/sysinfo"
)

# 遍历数组
for url in "${urls[@]}"; do
    curl="https://raw.githubusercontent.com/"$url"/main/go.mod"
    durl="https://raw.githubusercontent.com/"$url"/master/go.mod"
    echo $url
    # 使用curl获取URL内容，并打印出来
    # echo $curl"\n"
    curl -s $curl
    echo ""
    # echo $durl"\n"
    curl -s $durl
    echo "----------------------------------"
done
