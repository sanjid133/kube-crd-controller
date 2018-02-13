# kube-crd-controller
Practicing kubernetes CRD


* Create a SecDb type crd, which contains the list of secrets with their data
* Added a controller, which watches that crd
* If an item is added to a "secdb" then new secret will be created within a  namespace retrieve from "secdb"
* secrets of a type can be deleted and updated also by update that crd info.


