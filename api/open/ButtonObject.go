package open

type ButtonObject struct {
  Name        string `json:"name"`            // 菜单标题，一级菜单不超过4个汉字，子菜单不超过12个汉字
  ActionType  string `json:"action_type"`     // 菜单类型： out——事件型菜单； link——链接型菜单； tel——点击拨打电话
  ActionParam string `json:"action_param"`    // 当actionType为link时，该参数为详细链接； 当actionType为out时，该参数为用户自定义参数； 当actionType为tel时，该参数为电话号码。 该参数最长255个字符，不允许冒号等特殊字符
  SubButton   []SubButton `json:"sub_button"` // 二级菜单数组，个数应为1~5个
}

type SubButton struct {
  Name        string `json:"name"`            // 菜单标题，一级菜单不超过4个汉字，子菜单不超过12个汉字
  ActionType  string `json:"action_type"`     // 菜单类型： out——事件型菜单； link——链接型菜单； tel——点击拨打电话
  ActionParam string `json:"action_param"`    // 当actionType为link时，该参数为详细链接； 当actionType为out时，该参数为用户自定义参数； 当actionType为tel时，该参数为电话号码。 该参数最长255个字符，不允许冒号等特殊字符
}
