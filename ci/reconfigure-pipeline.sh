#!/bin/bash -ex
cd $(dirname $0)/..

flyrc_target=$1
if [ -z $flyrc_target ]; then
  echo "No target passed, using 'grootfs-ci'"
  flyrc_target="grootfs-ci"
fi

main() {
  check_fly_alias_exists
  sync_and_login

  pipeline_name="grootfs-diagnostics"
  vars_name="gcp"
  pipeline_file="pipeline.yml"

  set_pipeline
  expose_pipeline
}

set_pipeline() {
  fly --target="$flyrc_target" set-pipeline --pipeline=$pipeline_name \
    --config=ci/${pipeline_file} \
    --var slack-alert-url="$(lpass show 'Shared-Garden/grootfs-pivotal-slack-hook' --url)" \
    --var aws-access-key-id="$(lpass show "Shared-Garden/grootfs-deployments/grootfs-dstate-reports-s3-user" --username)" \
    --var aws-secret-access-key="$(lpass show "Shared-Garden/grootfs-deployments/grootfs-dstate-reports-s3-user" --password)" \
    --var datadog-api-key="$(lpass show 'Shared-Garden/grootfs-deployments/datadog-api-keys' --username)" \
    --var datadog-application-key="$(lpass show 'Shared-Garden/grootfs-deployments/datadog-api-keys' --password)" \
    --var bosh-target="https://10.0.0.6" \
    --var bosh-username="$(lpass show 'Shared-Garden/grootfs-deployments\thanos/bosh-director' --username)" \
    --var bosh-password="$(lpass show 'Shared-Garden/grootfs-deployments\thanos/bosh-director' --password)" \
    --var bosh-certificates="$(lpass show 'Shared-Garden/grootfs-deployments\thanos/certificates' --notes)" \
    --var thanos-cf-diego-certs="$(lpass show 'Shared-Garden/grootfs-deployments\thanos/cf-diego-certs' --notes)"
}

expose_pipeline() {
  fly --target="$flyrc_target" expose-pipeline --pipeline="$pipeline_name"
}

check_fly_alias_exists() {
  set +e
  grep $flyrc_target ~/.flyrc > /dev/null 2>&1
  if [ $? -ne 0 ]; then
    echo "Please ensure $flyrc_target exists in ~/.flyrc and that you have run fly login"
    exit 1
  fi
  set -e
}

sync_and_login() {
  fly -t $flyrc_target sync

  set +e
  fly -t $flyrc_target containers > /dev/null 2>&1
  if [[ $? -ne 0 ]]; then
    fly -t $flyrc_target login
  fi
  set -e
}

main
