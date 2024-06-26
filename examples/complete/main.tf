// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

module "private_ca" {
  source  = "terraform.registry.launch.nttdata.com/module_primitive/private_ca/aws"
  version = "~> 1.0"

  logical_product_family  = "demo"
  logical_product_service = "20534"

  tags = var.tags
}

module "private_cert" {
  source = "../.."

  private_ca_arn            = module.private_ca.private_ca_arn
  domain_name               = var.domain_name
  subject_alternative_names = var.subject_alternative_names

}
