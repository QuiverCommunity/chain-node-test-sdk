# Introduction

This SDK is for enhancing cosmos sdk projects which does not have integration test yet to have powerful integration test supported by our community.

# Features

- story test configuration on config.yml file.
node, cli_path, chain name, initial_block_height, stories_dir are configurable 
- automatic account creation from story's accounts section
- automatic running of transactions on provided blocks
- custom action registration running

# How to run this

Initialize node configuration and run node by running `sh ./test/test-env-setup.sh`.
Run integration test using `make integration_test`

# How to integrate this to your project

