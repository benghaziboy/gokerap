VAGRANTFILE_API_VERSION = "2"
Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  config.vm.box = "ubuntu/trusty64"
  config.vm.network "private_network", ip: "192.168.13.81"
  config.vm.network "forwarded_port", guest: 8200, host: 8200
  config.vm.provision "shell",
    inline: "apt-get -y update && apt-get -y install docker.io"
  config.vm.provision "shell",
    inline: "curl -L https://github.com/docker/fig/releases/download/0.5.2/linux > /usr/local/bin/fig && chmod +x /usr/local/bin/fig"
end
