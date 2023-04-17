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

set('dotenv', '{{current_path}}/.env');

// SERVERS
host('test')
    ->set('labels', ['stage' => 'development'])
    ->set('branch', 'develop')
    ->setHostname(getenv('DEP_HOST'))
    ->setRemoteUser('deployer')
    ->setDeployPath(getenv('DEP_DEPLOYPATH'))
    ->setIdentityFile('~/.ssh/id_rsa_home');

host('prod')
    ->set('labels', ['stage' => 'production'])
    ->set('branch', 'master')
    ->setHostname(getenv('DEP_HOST'))
    ->setRemoteUser('deployer')
    ->setDeployPath(getenv('DEP_DEPLOYPATH'))
    ->setIdentityFile('~/.ssh/id_rsa_home');
