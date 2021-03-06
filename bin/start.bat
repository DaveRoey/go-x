
set ETCD_HOSTS="192.168.1.4:12379, 192.168.1.4:22379, 192.168.1.4:32379"

start base.exe -cfg config.toml --etcd-hosts=%ETCD_HOSTS% --log_dir=./log -stderrthreshold 0
start gateway.exe -cfg config.toml --etcd-hosts=%ETCD_HOSTS% --log_dir=./log -stderrthreshold 0
start room.exe -cfg config.toml --etcd-hosts=%ETCD_HOSTS% --log_dir=./log -stderrthreshold 0
start login.exe -cfg config.toml --etcd-hosts=%ETCD_HOSTS% --log_dir=./log -stderrthreshold 0
