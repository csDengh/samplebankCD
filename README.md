# samplebankCD

代码仓库：
  push or pr 触发 CI：test->build and push image->send change to this CD.
  
配置仓库：
  receive ci change->change manifest

argocd:
  pull 配置仓库进行sync

manifest 的资源是自定义的，简化上层的配置，cr由controller进行拉起，发生变化，进行reconcile
