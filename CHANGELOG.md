# [1.1.0](https://github.com/veepee-oss/terraform-provider-awx/compare/v1.0.3...v1.1.0) (2025-11-03)


### Bug Fixes

* **goawx:** use latest version available ([89d069c](https://github.com/veepee-oss/terraform-provider-awx/commit/89d069ca81804a48f659f5d3187022f9f114e2b5))
* **organization_instance_group:** use correct input field ([7ee0c5d](https://github.com/veepee-oss/terraform-provider-awx/commit/7ee0c5d23667d8bf97ffbf1614c5bfd31ef6a875))
* **provider:** typo ([6ef4e74](https://github.com/veepee-oss/terraform-provider-awx/commit/6ef4e747a6fa4e7174eb53766e15e7d8f4d356e5))
* **resourde:** rename organization_instance_group ([5fda8ef](https://github.com/veepee-oss/terraform-provider-awx/commit/5fda8ef1f219597672909a51ccbde3e4663c4030))


### Features

* **ressources:** manage organization instance group endpoint ([55637fd](https://github.com/veepee-oss/terraform-provider-awx/commit/55637fd4d6707e40a33d7770a75c49298d43a0cf))

## [1.0.3](https://github.com/veepee-oss/terraform-provider-awx/compare/v1.0.2...v1.0.3) (2025-07-30)


### Bug Fixes

* **goreleaser:** move to configuration v2 and remove changelog ([7543a26](https://github.com/veepee-oss/terraform-provider-awx/commit/7543a2619bee29e0be2f2b0304d0106ccd2aa27f))

## [1.0.2](https://github.com/veepee-oss/terraform-provider-awx/compare/v1.0.1...v1.0.2) (2025-07-30)


### Bug Fixes

* **release:** use new args to clean dist instead of deprecated --rm-dist ([4ce36ba](https://github.com/veepee-oss/terraform-provider-awx/commit/4ce36ba4206224d14e381956a1f0d650c650b497))

## [1.0.1](https://github.com/veepee-oss/terraform-provider-awx/compare/v1.0.0...v1.0.1) (2025-07-30)


### Bug Fixes

* **credential:** allow to define galaxy credential type id ([f9591ec](https://github.com/veepee-oss/terraform-provider-awx/commit/f9591ec681e1a69cc8ff36d0fba94d0594caaaee))
* **credential:** move credential_type_id to default ([b1bf2df](https://github.com/veepee-oss/terraform-provider-awx/commit/b1bf2dff0dbebfbe91a164d5c9a234b99c9075f3))

# 1.0.0 (2023-11-14)


### Bug Fixes

* add missing unified_job_template_id ([72d1deb](https://github.com/veepee-oss/terraform-provider-awx/commit/72d1deb810d8618158bb48cea924959961495163))
* add missing workflow_job_template_id on workflow_job_template_node_* ([ca6c20f](https://github.com/veepee-oss/terraform-provider-awx/commit/ca6c20fc412d8a9dff511b8baffc20412d77a6c8))
* add schedule inventory parameter support ([5e691ac](https://github.com/veepee-oss/terraform-provider-awx/commit/5e691ac67f0e28337688928a96d6a3f1b0a7376a))
* do not provide local_path for project if the scm_type is git ([#13](https://github.com/veepee-oss/terraform-provider-awx/issues/13)) ([b4ab7dc](https://github.com/veepee-oss/terraform-provider-awx/commit/b4ab7dc51306507bd71ef61b611782567bc0c0bb))
* fix some func names after upgrading goawx dep ([999e70d](https://github.com/veepee-oss/terraform-provider-awx/commit/999e70ddbdcdc3ca758b85e9c6a4eea3b3859689))
* goawx version for always node type ([#7](https://github.com/veepee-oss/terraform-provider-awx/issues/7)) ([bfe6ea8](https://github.com/veepee-oss/terraform-provider-awx/commit/bfe6ea8d2245836a5b2584b4d471ca911d1b4626))
* make a new release ([be91fb4](https://github.com/veepee-oss/terraform-provider-awx/commit/be91fb4577e932ffee1019efb70620479d6089fd))
* notification template notification configuration is a json ([09787ef](https://github.com/veepee-oss/terraform-provider-awx/commit/09787ef93e745a0049970f5fcd134f5ab5a7f6f5))
* notification_configuration is a string ([f10fb3b](https://github.com/veepee-oss/terraform-provider-awx/commit/f10fb3ba03deca84d3169bc2eac0b01503c438f8))
* notification_template schema ([4b28594](https://github.com/veepee-oss/terraform-provider-awx/commit/4b2859405fc56bb7a09320f826862cbaa05a6d32))
* publish for all os and arch ([7a3cd45](https://github.com/veepee-oss/terraform-provider-awx/commit/7a3cd4552b44246377a00a185dbde48b45ce07dc))
* remove success_nodes from workflow_job_template_node ([5957d52](https://github.com/veepee-oss/terraform-provider-awx/commit/5957d526162cb9ebdcc237058af9c29825b0315b))
* rrule is a string ([26f9404](https://github.com/veepee-oss/terraform-provider-awx/commit/26f9404c64497388ee040a97ef8b6e6271827f15))
* schedule Optional field missing ([acc2538](https://github.com/veepee-oss/terraform-provider-awx/commit/acc2538436c739549e87b7686d54880d013708cb))
* some fixes ([196711f](https://github.com/veepee-oss/terraform-provider-awx/commit/196711fa77569ec58bf54716ac5f81e736278f77))
* some fixes on notification_template resource ([3cd1a59](https://github.com/veepee-oss/terraform-provider-awx/commit/3cd1a592ad1c3baed7a237aa228645a90cb790cb))
* upgrade goawx dep ([ba2ea50](https://github.com/veepee-oss/terraform-provider-awx/commit/ba2ea509f164f7dad4f5477d6d58a40a798c0022))
* upgrade goawx dep ([50447a2](https://github.com/veepee-oss/terraform-provider-awx/commit/50447a2ebf2a0fb2862f2749a6aaa7ec58fed0e7))
* use roondar/goawx v1.1.1 ([b28f568](https://github.com/veepee-oss/terraform-provider-awx/commit/b28f5684e8ff14fffbc999e9d3cbb62593aa2c9d))
* when using insecure connection the PROXY_HTTPS env var was ignored ([#12](https://github.com/veepee-oss/terraform-provider-awx/issues/12)) ([e457deb](https://github.com/veepee-oss/terraform-provider-awx/commit/e457deb4644f82e4c4e3af27e07df7ba565cbbaa))
* workflow job template & schedule inventory option default value ([#2](https://github.com/veepee-oss/terraform-provider-awx/issues/2)) ([6869420](https://github.com/veepee-oss/terraform-provider-awx/commit/6869420d6b87a70922c915d1012ebd15156a277a))
* workflow_job_template_node_id is required ([a283788](https://github.com/veepee-oss/terraform-provider-awx/commit/a2837882f3b4f5d59c8d65bb309bfd54ed97940c))


### Features

* add AWX token authentication ([#15](https://github.com/veepee-oss/terraform-provider-awx/issues/15)) ([55b7d41](https://github.com/veepee-oss/terraform-provider-awx/commit/55b7d41579f79d7fbb1aa61bd243405a81815748))
* add awx_schedule and awx_workflow_job_template_schedule resources ([af3ec75](https://github.com/veepee-oss/terraform-provider-awx/commit/af3ec75da0893d7d964a63777be60cfc4508dd41))
* Add default environment for Organizations ([6f8a2aa](https://github.com/veepee-oss/terraform-provider-awx/commit/6f8a2aa786cca31eeba249e093add818abcbb0a2))
* Add doc to gitlab ([68f2d76](https://github.com/veepee-oss/terraform-provider-awx/commit/68f2d763fb6f9ecf50b08a05688f374b4803860e))
* Add Gitlab credential ([929b5fd](https://github.com/veepee-oss/terraform-provider-awx/commit/929b5fd43d970f170c08779720fb26133b869627))
* add notification_template resource ([9c5b488](https://github.com/veepee-oss/terraform-provider-awx/commit/9c5b4885dfcd068b7dbac89567067c606b73fa6c))
* Add organization role ([2d19fb2](https://github.com/veepee-oss/terraform-provider-awx/commit/2d19fb23abf7a14d6148e0bac80c84a6ecaa0fe0))
* Add organization role data source ([4bc4065](https://github.com/veepee-oss/terraform-provider-awx/commit/4bc40653f96d92c47d0e0f5fb53d4172661491f4))
* Add Organizations GalaxyCredentials ([56c7971](https://github.com/veepee-oss/terraform-provider-awx/commit/56c797102713d12df04852312b7b7b4bac9f7789))
* Add resource credential Ansible Galaxy ([fe00422](https://github.com/veepee-oss/terraform-provider-awx/commit/fe00422223ec3ab051b3078125794efbe25c1381))
* Add resource settings ([c3e2646](https://github.com/veepee-oss/terraform-provider-awx/commit/c3e2646820ab2bca3a310a32be98ca077c07c47e))
* add resources awx_job_template_notification_template_success awx_job_template_notification_template_error awx_job_template_notification_template_started ([24b69c5](https://github.com/veepee-oss/terraform-provider-awx/commit/24b69c5ded4c0fbba366637c0e423e0fc07679e6))
* Add setting resource ([b1b1a24](https://github.com/veepee-oss/terraform-provider-awx/commit/b1b1a2403887599bd451e54094a48c7e728aa8da))
* Add source project and source path for inventory ([d881d7f](https://github.com/veepee-oss/terraform-provider-awx/commit/d881d7f78aa903473b6d7a45620dca97236f59a3))
* Add user resource ([3af34a2](https://github.com/veepee-oss/terraform-provider-awx/commit/3af34a23a4d8f35c99f273a720012c1c4a19dfed))
* adds the possibility to use source_id inside resource_inventory_source ([#20](https://github.com/veepee-oss/terraform-provider-awx/issues/20)) ([6891c9e](https://github.com/veepee-oss/terraform-provider-awx/commit/6891c9eb98b8ca916746af08c640520f07d29dda))
* enable insecure https connection to AWX ([#84](https://github.com/veepee-oss/terraform-provider-awx/issues/84)) ([616e88d](https://github.com/veepee-oss/terraform-provider-awx/commit/616e88da2be22516413ad5ffa8a48152a2095050))
* fetch upstream ([8cc9cb0](https://github.com/veepee-oss/terraform-provider-awx/commit/8cc9cb0f160b779e02f17fab10c03d4cb7ec54b9)), closes [#16](https://github.com/veepee-oss/terraform-provider-awx/issues/16)
* organizations data source ([#4](https://github.com/veepee-oss/terraform-provider-awx/issues/4)) ([ad61e88](https://github.com/veepee-oss/terraform-provider-awx/commit/ad61e88a638b94eda2c306a0d9f610d65508d17f))
* release ([ab77b3a](https://github.com/veepee-oss/terraform-provider-awx/commit/ab77b3afef3e78c20abc71d1f88e4b8ef2f8ef05))
* support execution environments ([#1](https://github.com/veepee-oss/terraform-provider-awx/issues/1)) ([0791c09](https://github.com/veepee-oss/terraform-provider-awx/commit/0791c09cb85783e7433f8e4ea80cfa9d7911af32))
* Update goawx v0.18.0 ([0a753d8](https://github.com/veepee-oss/terraform-provider-awx/commit/0a753d8a997cc1c803266200948dad06ceef1ee3))
* upgrade goawx & support for job_slice_count ([f7b6198](https://github.com/veepee-oss/terraform-provider-awx/commit/f7b61982af53ecc64c4e38ddab112f8b56a64857))
* upgrade goawx lib to 0.14.1 ([#22](https://github.com/veepee-oss/terraform-provider-awx/issues/22)) ([3193f56](https://github.com/veepee-oss/terraform-provider-awx/commit/3193f56a55ac96103f0b2a4f355f0cd723116f86))
* Use roondar/goawx v1.1.0 ([7288f19](https://github.com/veepee-oss/terraform-provider-awx/commit/7288f19b1fc6cfc4e413041edb189db14e32b216))
* workflow job template notifications ([#3](https://github.com/veepee-oss/terraform-provider-awx/issues/3)) ([00db915](https://github.com/veepee-oss/terraform-provider-awx/commit/00db9157df52d9fb4431db6f53ac5aa8038bad44))

# [1.3.0](https://github.com/veepee-oss/terraform-provider-awx/compare/v1.2.1...v1.3.0) (2023-11-14)


### Features

* upgrade goawx & support for job_slice_count ([f7b6198](https://github.com/veepee-oss/terraform-provider-awx/commit/f7b61982af53ecc64c4e38ddab112f8b56a64857))

## [1.2.1](https://github.com/roondar/terraform-provider-awx/compare/v1.2.0...v1.2.1) (2023-04-05)


### Bug Fixes

* use roondar/goawx v1.1.1 ([b28f568](https://github.com/roondar/terraform-provider-awx/commit/b28f5684e8ff14fffbc999e9d3cbb62593aa2c9d))

# [1.2.0](https://github.com/roondar/terraform-provider-awx/compare/v1.1.0...v1.2.0) (2023-04-05)


### Features

* Use roondar/goawx v1.1.0 ([7288f19](https://github.com/roondar/terraform-provider-awx/commit/7288f19b1fc6cfc4e413041edb189db14e32b216))

# [1.1.0](https://github.com/roondar/terraform-provider-awx/compare/v1.0.0...v1.1.0) (2023-04-04)


### Features

* Update goawx v0.18.0 ([0a753d8](https://github.com/roondar/terraform-provider-awx/commit/0a753d8a997cc1c803266200948dad06ceef1ee3))

# 1.0.0 (2023-04-04)


### Bug Fixes

* add missing unified_job_template_id ([72d1deb](https://github.com/roondar/terraform-provider-awx/commit/72d1deb810d8618158bb48cea924959961495163))
* add missing workflow_job_template_id on workflow_job_template_node_* ([ca6c20f](https://github.com/roondar/terraform-provider-awx/commit/ca6c20fc412d8a9dff511b8baffc20412d77a6c8))
* add schedule inventory parameter support ([5e691ac](https://github.com/roondar/terraform-provider-awx/commit/5e691ac67f0e28337688928a96d6a3f1b0a7376a))
* do not provide local_path for project if the scm_type is git ([#13](https://github.com/roondar/terraform-provider-awx/issues/13)) ([b4ab7dc](https://github.com/roondar/terraform-provider-awx/commit/b4ab7dc51306507bd71ef61b611782567bc0c0bb))
* fix some func names after upgrading goawx dep ([999e70d](https://github.com/roondar/terraform-provider-awx/commit/999e70ddbdcdc3ca758b85e9c6a4eea3b3859689))
* goawx version for always node type ([#7](https://github.com/roondar/terraform-provider-awx/issues/7)) ([bfe6ea8](https://github.com/roondar/terraform-provider-awx/commit/bfe6ea8d2245836a5b2584b4d471ca911d1b4626))
* make a new release ([be91fb4](https://github.com/roondar/terraform-provider-awx/commit/be91fb4577e932ffee1019efb70620479d6089fd))
* notification template notification configuration is a json ([09787ef](https://github.com/roondar/terraform-provider-awx/commit/09787ef93e745a0049970f5fcd134f5ab5a7f6f5))
* notification_configuration is a string ([f10fb3b](https://github.com/roondar/terraform-provider-awx/commit/f10fb3ba03deca84d3169bc2eac0b01503c438f8))
* notification_template schema ([4b28594](https://github.com/roondar/terraform-provider-awx/commit/4b2859405fc56bb7a09320f826862cbaa05a6d32))
* publish for all os and arch ([7a3cd45](https://github.com/roondar/terraform-provider-awx/commit/7a3cd4552b44246377a00a185dbde48b45ce07dc))
* remove success_nodes from workflow_job_template_node ([5957d52](https://github.com/roondar/terraform-provider-awx/commit/5957d526162cb9ebdcc237058af9c29825b0315b))
* rrule is a string ([26f9404](https://github.com/roondar/terraform-provider-awx/commit/26f9404c64497388ee040a97ef8b6e6271827f15))
* schedule Optional field missing ([acc2538](https://github.com/roondar/terraform-provider-awx/commit/acc2538436c739549e87b7686d54880d013708cb))
* some fixes ([196711f](https://github.com/roondar/terraform-provider-awx/commit/196711fa77569ec58bf54716ac5f81e736278f77))
* some fixes on notification_template resource ([3cd1a59](https://github.com/roondar/terraform-provider-awx/commit/3cd1a592ad1c3baed7a237aa228645a90cb790cb))
* upgrade goawx dep ([ba2ea50](https://github.com/roondar/terraform-provider-awx/commit/ba2ea509f164f7dad4f5477d6d58a40a798c0022))
* upgrade goawx dep ([50447a2](https://github.com/roondar/terraform-provider-awx/commit/50447a2ebf2a0fb2862f2749a6aaa7ec58fed0e7))
* when using insecure connection the PROXY_HTTPS env var was ignored ([#12](https://github.com/roondar/terraform-provider-awx/issues/12)) ([e457deb](https://github.com/roondar/terraform-provider-awx/commit/e457deb4644f82e4c4e3af27e07df7ba565cbbaa))
* workflow job template & schedule inventory option default value ([#2](https://github.com/roondar/terraform-provider-awx/issues/2)) ([6869420](https://github.com/roondar/terraform-provider-awx/commit/6869420d6b87a70922c915d1012ebd15156a277a))
* workflow_job_template_node_id is required ([a283788](https://github.com/roondar/terraform-provider-awx/commit/a2837882f3b4f5d59c8d65bb309bfd54ed97940c))


### Features

* add AWX token authentication ([#15](https://github.com/roondar/terraform-provider-awx/issues/15)) ([55b7d41](https://github.com/roondar/terraform-provider-awx/commit/55b7d41579f79d7fbb1aa61bd243405a81815748))
* add awx_schedule and awx_workflow_job_template_schedule resources ([af3ec75](https://github.com/roondar/terraform-provider-awx/commit/af3ec75da0893d7d964a63777be60cfc4508dd41))
* Add default environment for Organizations ([6f8a2aa](https://github.com/roondar/terraform-provider-awx/commit/6f8a2aa786cca31eeba249e093add818abcbb0a2))
* Add doc to gitlab ([68f2d76](https://github.com/roondar/terraform-provider-awx/commit/68f2d763fb6f9ecf50b08a05688f374b4803860e))
* Add Gitlab credential ([929b5fd](https://github.com/roondar/terraform-provider-awx/commit/929b5fd43d970f170c08779720fb26133b869627))
* add notification_template resource ([9c5b488](https://github.com/roondar/terraform-provider-awx/commit/9c5b4885dfcd068b7dbac89567067c606b73fa6c))
* Add organization role ([2d19fb2](https://github.com/roondar/terraform-provider-awx/commit/2d19fb23abf7a14d6148e0bac80c84a6ecaa0fe0))
* Add organization role data source ([4bc4065](https://github.com/roondar/terraform-provider-awx/commit/4bc40653f96d92c47d0e0f5fb53d4172661491f4))
* Add Organizations GalaxyCredentials ([56c7971](https://github.com/roondar/terraform-provider-awx/commit/56c797102713d12df04852312b7b7b4bac9f7789))
* Add resource credential Ansible Galaxy ([fe00422](https://github.com/roondar/terraform-provider-awx/commit/fe00422223ec3ab051b3078125794efbe25c1381))
* Add resource settings ([c3e2646](https://github.com/roondar/terraform-provider-awx/commit/c3e2646820ab2bca3a310a32be98ca077c07c47e))
* add resources awx_job_template_notification_template_success awx_job_template_notification_template_error awx_job_template_notification_template_started ([24b69c5](https://github.com/roondar/terraform-provider-awx/commit/24b69c5ded4c0fbba366637c0e423e0fc07679e6))
* Add setting resource ([b1b1a24](https://github.com/roondar/terraform-provider-awx/commit/b1b1a2403887599bd451e54094a48c7e728aa8da))
* Add source project and source path for inventory ([d881d7f](https://github.com/roondar/terraform-provider-awx/commit/d881d7f78aa903473b6d7a45620dca97236f59a3))
* Add user resource ([3af34a2](https://github.com/roondar/terraform-provider-awx/commit/3af34a23a4d8f35c99f273a720012c1c4a19dfed))
* adds the possibility to use source_id inside resource_inventory_source ([#20](https://github.com/roondar/terraform-provider-awx/issues/20)) ([6891c9e](https://github.com/roondar/terraform-provider-awx/commit/6891c9eb98b8ca916746af08c640520f07d29dda))
* enable insecure https connection to AWX ([#84](https://github.com/roondar/terraform-provider-awx/issues/84)) ([616e88d](https://github.com/roondar/terraform-provider-awx/commit/616e88da2be22516413ad5ffa8a48152a2095050))
* fetch upstream ([8cc9cb0](https://github.com/roondar/terraform-provider-awx/commit/8cc9cb0f160b779e02f17fab10c03d4cb7ec54b9)), closes [#16](https://github.com/roondar/terraform-provider-awx/issues/16)
* organizations data source ([#4](https://github.com/roondar/terraform-provider-awx/issues/4)) ([ad61e88](https://github.com/roondar/terraform-provider-awx/commit/ad61e88a638b94eda2c306a0d9f610d65508d17f))
* support execution environments ([#1](https://github.com/roondar/terraform-provider-awx/issues/1)) ([0791c09](https://github.com/roondar/terraform-provider-awx/commit/0791c09cb85783e7433f8e4ea80cfa9d7911af32))
* upgrade goawx lib to 0.14.1 ([#22](https://github.com/roondar/terraform-provider-awx/issues/22)) ([3193f56](https://github.com/roondar/terraform-provider-awx/commit/3193f56a55ac96103f0b2a4f355f0cd723116f86))
* workflow job template notifications ([#3](https://github.com/roondar/terraform-provider-awx/issues/3)) ([00db915](https://github.com/roondar/terraform-provider-awx/commit/00db9157df52d9fb4431db6f53ac5aa8038bad44))

# [0.19.0](https://github.com/denouche/terraform-provider-awx/compare/v0.18.0...v0.19.0) (2022-11-14)


### Features

* Add organization role data source ([4bc4065](https://github.com/denouche/terraform-provider-awx/commit/4bc40653f96d92c47d0e0f5fb53d4172661491f4))
* Add setting resource ([b1b1a24](https://github.com/denouche/terraform-provider-awx/commit/b1b1a2403887599bd451e54094a48c7e728aa8da))
* fetch upstream ([8cc9cb0](https://github.com/denouche/terraform-provider-awx/commit/8cc9cb0f160b779e02f17fab10c03d4cb7ec54b9)), closes [#16](https://github.com/denouche/terraform-provider-awx/issues/16)

# [0.18.0](https://github.com/denouche/terraform-provider-awx/compare/v0.17.0...v0.18.0) (2022-11-14)


### Features

* add AWX token authentication ([#15](https://github.com/denouche/terraform-provider-awx/issues/15)) ([55b7d41](https://github.com/denouche/terraform-provider-awx/commit/55b7d41579f79d7fbb1aa61bd243405a81815748))

# [0.17.0](https://github.com/denouche/terraform-provider-awx/compare/v0.16.0...v0.17.0) (2022-10-31)


### Features

* adds the possibility to use source_id inside resource_inventory_source ([#20](https://github.com/denouche/terraform-provider-awx/issues/20)) ([6891c9e](https://github.com/denouche/terraform-provider-awx/commit/6891c9eb98b8ca916746af08c640520f07d29dda))

# [0.16.0](https://github.com/denouche/terraform-provider-awx/compare/v0.15.6...v0.16.0) (2022-10-25)


### Features

* upgrade goawx lib to 0.14.1 ([#22](https://github.com/denouche/terraform-provider-awx/issues/22)) ([3193f56](https://github.com/denouche/terraform-provider-awx/commit/3193f56a55ac96103f0b2a4f355f0cd723116f86))

## [0.15.6](https://github.com/denouche/terraform-provider-awx/compare/v0.15.5...v0.15.6) (2022-07-22)


### Bug Fixes

* when using insecure connection the PROXY_HTTPS env var was ignored ([#12](https://github.com/denouche/terraform-provider-awx/issues/12)) ([e457deb](https://github.com/denouche/terraform-provider-awx/commit/e457deb4644f82e4c4e3af27e07df7ba565cbbaa))

## [0.15.5](https://github.com/denouche/terraform-provider-awx/compare/v0.15.4...v0.15.5) (2022-07-22)


### Bug Fixes

* fix some func names after upgrading goawx dep ([999e70d](https://github.com/denouche/terraform-provider-awx/commit/999e70ddbdcdc3ca758b85e9c6a4eea3b3859689))

## [0.15.4](https://github.com/denouche/terraform-provider-awx/compare/v0.15.3...v0.15.4) (2022-07-22)

## [0.15.3](https://github.com/denouche/terraform-provider-awx/compare/v0.15.2...v0.15.3) (2022-07-19)


### Bug Fixes

* do not provide local_path for project if the scm_type is git ([#13](https://github.com/denouche/terraform-provider-awx/issues/13)) ([b4ab7dc](https://github.com/denouche/terraform-provider-awx/commit/b4ab7dc51306507bd71ef61b611782567bc0c0bb))

## [0.15.2](https://github.com/denouche/terraform-provider-awx/compare/v0.15.1...v0.15.2) (2022-07-01)


### Bug Fixes

* make a new release ([be91fb4](https://github.com/denouche/terraform-provider-awx/commit/be91fb4577e932ffee1019efb70620479d6089fd))

## [0.15.1](https://github.com/denouche/terraform-provider-awx/compare/v0.15.0...v0.15.1) (2022-07-01)


### Bug Fixes

* goawx version for always node type ([#7](https://github.com/denouche/terraform-provider-awx/issues/7)) ([bfe6ea8](https://github.com/denouche/terraform-provider-awx/commit/bfe6ea8d2245836a5b2584b4d471ca911d1b4626))

# [0.15.0](https://github.com/denouche/terraform-provider-awx/compare/v0.14.0...v0.15.0) (2022-05-11)


### Features

* organizations data source ([#4](https://github.com/denouche/terraform-provider-awx/issues/4)) ([ad61e88](https://github.com/denouche/terraform-provider-awx/commit/ad61e88a638b94eda2c306a0d9f610d65508d17f))

# [0.14.0](https://github.com/denouche/terraform-provider-awx/compare/v0.13.1...v0.14.0) (2022-04-21)


### Features

* workflow job template notifications ([#3](https://github.com/denouche/terraform-provider-awx/issues/3)) ([00db915](https://github.com/denouche/terraform-provider-awx/commit/00db9157df52d9fb4431db6f53ac5aa8038bad44))

## [0.13.1](https://github.com/denouche/terraform-provider-awx/compare/v0.13.0...v0.13.1) (2022-04-20)


### Bug Fixes

* workflow job template & schedule inventory option default value ([#2](https://github.com/denouche/terraform-provider-awx/issues/2)) ([6869420](https://github.com/denouche/terraform-provider-awx/commit/6869420d6b87a70922c915d1012ebd15156a277a))

# [0.13.0](https://github.com/denouche/terraform-provider-awx/compare/v0.12.3...v0.13.0) (2022-04-20)


### Features

* support execution environments ([#1](https://github.com/denouche/terraform-provider-awx/issues/1)) ([0791c09](https://github.com/denouche/terraform-provider-awx/commit/0791c09cb85783e7433f8e4ea80cfa9d7911af32))

## [0.12.3](https://github.com/denouche/terraform-provider-awx/compare/v0.12.2...v0.12.3) (2022-04-19)


### Bug Fixes

* publish for all os and arch ([7a3cd45](https://github.com/denouche/terraform-provider-awx/commit/7a3cd4552b44246377a00a185dbde48b45ce07dc))

## [0.12.2](https://github.com/denouche/terraform-provider-awx/compare/v0.12.1...v0.12.2) (2022-01-05)


### Bug Fixes

* upgrade goawx dep ([ba2ea50](https://github.com/denouche/terraform-provider-awx/commit/ba2ea509f164f7dad4f5477d6d58a40a798c0022))

## [0.12.1](https://github.com/denouche/terraform-provider-awx/compare/v0.12.0...v0.12.1) (2022-01-05)


### Bug Fixes

* upgrade goawx dep ([50447a2](https://github.com/denouche/terraform-provider-awx/commit/50447a2ebf2a0fb2862f2749a6aaa7ec58fed0e7))

# [0.12.0](https://github.com/denouche/terraform-provider-awx/compare/v0.11.4...v0.12.0) (2022-01-05)


### Features

* add resources awx_job_template_notification_template_success awx_job_template_notification_template_error awx_job_template_notification_template_started ([24b69c5](https://github.com/denouche/terraform-provider-awx/commit/24b69c5ded4c0fbba366637c0e423e0fc07679e6))

## [0.11.4](https://github.com/denouche/terraform-provider-awx/compare/v0.11.3...v0.11.4) (2021-12-24)


### Bug Fixes

* notification template notification configuration is a json ([09787ef](https://github.com/denouche/terraform-provider-awx/commit/09787ef93e745a0049970f5fcd134f5ab5a7f6f5))

## [0.11.3](https://github.com/denouche/terraform-provider-awx/compare/v0.11.2...v0.11.3) (2021-12-24)


### Bug Fixes

* notification_configuration is a string ([f10fb3b](https://github.com/denouche/terraform-provider-awx/commit/f10fb3ba03deca84d3169bc2eac0b01503c438f8))

## [0.11.2](https://github.com/denouche/terraform-provider-awx/compare/v0.11.1...v0.11.2) (2021-12-24)


### Bug Fixes

* some fixes on notification_template resource ([3cd1a59](https://github.com/denouche/terraform-provider-awx/commit/3cd1a592ad1c3baed7a237aa228645a90cb790cb))

## [0.11.1](https://github.com/denouche/terraform-provider-awx/compare/v0.11.0...v0.11.1) (2021-12-24)


### Bug Fixes

* notification_template schema ([4b28594](https://github.com/denouche/terraform-provider-awx/commit/4b2859405fc56bb7a09320f826862cbaa05a6d32))

# [0.11.0](https://github.com/denouche/terraform-provider-awx/compare/v0.10.7...v0.11.0) (2021-12-24)


### Features

* add notification_template resource ([9c5b488](https://github.com/denouche/terraform-provider-awx/commit/9c5b4885dfcd068b7dbac89567067c606b73fa6c))

## [0.10.7](https://github.com/denouche/terraform-provider-awx/compare/v0.10.6...v0.10.7) (2021-12-23)


### Bug Fixes

* add missing unified_job_template_id ([72d1deb](https://github.com/denouche/terraform-provider-awx/commit/72d1deb810d8618158bb48cea924959961495163))

## [0.10.6](https://github.com/denouche/terraform-provider-awx/compare/v0.10.5...v0.10.6) (2021-12-23)


### Bug Fixes

* add schedule inventory parameter support ([5e691ac](https://github.com/denouche/terraform-provider-awx/commit/5e691ac67f0e28337688928a96d6a3f1b0a7376a))

## [0.10.5](https://github.com/denouche/terraform-provider-awx/compare/v0.10.4...v0.10.5) (2021-12-23)

## [0.10.4](https://github.com/denouche/terraform-provider-awx/compare/v0.10.3...v0.10.4) (2021-12-23)

## [0.10.3](https://github.com/denouche/terraform-provider-awx/compare/v0.10.2...v0.10.3) (2021-12-23)
