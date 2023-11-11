package translate

func WeaponNameZh(weapon string) string {
	switch weapon {
	case "P2020":
		return "P2020手枪"
	case "RE45":
		return "RE45"
	case "Wingman":
		return "辅助手枪"

	case "R99":
		return "R99冲锋枪"
	case "Prowler":
		return "猎兽冲锋枪"
	case "Alternator":
		return "转换者冲锋枪"
	case "Car":
		return "CAR"
	case "Volt SMG":
		return "电能冲锋枪"

	case "VK-47":
		return "平行步枪"
	case "Hemlok":
		return "赫姆洛克步枪"
	case "R-301":
		return "R-301"
	case "Havoc":
		return "哈沃克步枪"
	case "Nemesis":
		return "复仇女神"

	case "Devotion":
		return "专注轻机枪"
	case "Spitfire":
		return "喷火轻机枪"
	case "L-Star":
		return "L-Star"
	case "Rampage":
		return "暴走"

	case "Mastiff":
		return "敖犬霰弹枪"
	case "EVA-8":
		return "EVA-8"
	case "Peacekeeper":
		return "和平捍卫者霰弹枪"
	case "Mozambique":
		return "莫三比克"

	case "Charge Rifle":
		return "充能步枪"
	case "Longbow":
		return "长弓"
	case "Sentinel":
		return "哨兵狙击步枪"
	case "Kraber":
		return "克雷贝尔狙击枪"

	case "G7":
		return "G7侦察枪"
	case "30-30":
		return "30-30"
	case "Triple Take":
		return "三重式狙击枪"
	case "Bocek":
		return "博塞克"

	default:
		return weapon
	}
}
