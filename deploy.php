<?php

namespace Deployer;

require 'recipe/common.php';

// SETTINGS
set('keep_releases', 5);
set('repository', 'git@github.com:dmsemenov/carRegionBot.git');
set('shared_files', [
    '.env',
    'docker-compose.yml'
]);

// SERVERS
host('test')
    ->set('labels', ['stage' => 'development'])
    ->set('branch', 'develop')
    ->setHostname('185.178.44.125')
    ->setRemoteUser('deployer')
    ->setDeployPath('/home/deployer/projects/carregionbot.1108811-cd43481.tw1.ru')
    ->setIdentityFile('~/.ssh/id_rsa_home');

host('prod')
    ->set('labels', ['stage' => 'production'])
    ->set('branch', 'master')
    ->setHostname('185.178.44.125')
    ->setRemoteUser('deployer')
    ->setDeployPath('/home/deployer/projects/carregionbot.1108811-cd43481.tw1.ru')
    ->setIdentityFile('~/.ssh/id_rsa_home');
