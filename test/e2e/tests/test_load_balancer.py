# Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License"). You may
# not use this file except in compliance with the License. A copy of the
# License is located at
#
# 	 http://aws.amazon.com/apache2.0/
#
# or in the "license" file accompanying this file. This file is distributed
# on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
# express or implied. See the License for the specific language governing
# permissions and limitations under the License.

"""Integration tests for the ELB LoadBalancer API.
"""

import pytest
import time
import logging

from acktest.resources import random_suffix_name
from acktest.k8s import resource as k8s

from e2e import service_marker, CRD_GROUP, CRD_VERSION, load_elbv2_resource
from e2e.replacement_values import REPLACEMENT_VALUES
from e2e.bootstrap_resources import get_bootstrap_resources
from e2e.tests.helper import ELBValidator

RESOURCE_PLURAL = "loadbalancers"

CREATE_WAIT_AFTER_SECONDS = 10
UPDATE_WAIT_AFTER_SECONDS = 10
DELETE_WAIT_AFTER_SECONDS = 10

@pytest.fixture(scope="module")
def simple_load_balancer(elbv2_client):

    resource_name = random_suffix_name("lb", 16)

    replacements = REPLACEMENT_VALUES.copy()
    replacements["LOAD_BALANCER_NAME"] = resource_name

    resource_data = load_elbv2_resource(
        "load_balancer",
        additional_replacements=replacements,
    )
    logging.debug(resource_data)

    # Create k8s resource
    ref = k8s.CustomResourceReference(
        CRD_GROUP, CRD_VERSION, RESOURCE_PLURAL,
        resource_name, namespace="default",
    )
    k8s.create_custom_resource(ref, resource_data)

    time.sleep(CREATE_WAIT_AFTER_SECONDS)
    cr = k8s.wait_resource_consumed_by_controller(ref)

    assert cr is not None
    assert k8s.get_resource_exists(ref)

    yield (ref, cr, resource_name)

    _, deleted = k8s.delete_custom_resource(
        ref,
        period_length=DELETE_WAIT_AFTER_SECONDS,
    )
    assert deleted

    time.sleep(DELETE_WAIT_AFTER_SECONDS)

    validator = ELBValidator(elbv2_client)
    assert not validator.load_balancer_exists(resource_name)

@service_marker
@pytest.mark.canary
class TestLoadBalancer:
    def test_create_delete(self, elbv2_client, simple_load_balancer):
        (ref, _, load_balancer) = simple_load_balancer
        assert load_balancer is not None

        validator = ELBValidator(elbv2_client)
        assert validator.load_balancer_exists(load_balancer)

        # Update settings
        """ updates = {
            "spec": {
                "settings": [
                    {
                        "name": "containerInsights",
                        "value": "enabled"
                    }
                ],
            },
        }
        k8s.patch_custom_resource(ref, updates)
        time.sleep(UPDATE_WAIT_AFTER_SECONDS)

        cs = validator.get_cluster(cluster_name)
        assert cs["clusters"][0]["settings"][0]["name"] == "containerInsights"
        assert cs["clusters"][0]["settings"][0]["value"] == "enabled" """