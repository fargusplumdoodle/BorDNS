
helm uninstall testbordns 
helm install testbordns . \
	--set etcd.storage.nfs.server=10.0.1.30 \
	--set etcd.storage.nfs.path=/mnt/horus/k8-data/bordns \
