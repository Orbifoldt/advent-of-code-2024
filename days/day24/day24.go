package day24

import (
	"advent-of-code-2024/util"
	"fmt"
	"slices"
	"strings"
)

func SolvePart1(useRealInput bool) (int, error) {
	board, err := parseInput(useRealInput)
	if err != nil {
		return 0, err
	}

	outputWireLabels := make([]string, 0)
	for label := range board.wires {
		if label[0] == 'z' {
			outputWireLabels = append(outputWireLabels, label)
		}
	}
	slices.Sort(outputWireLabels)
	slices.Reverse(outputWireLabels)

	output := 0
	for _, label := range outputWireLabels {
		var value int
		if board.getWireValue(label) {
			value = 1
		} else {
			value = 0
		}
		// fmt.Printf("%s=%d\n", label, value)
		output = (output << 1) + value
	}

	return output, nil
}

type Operator int

const (
	AND Operator = iota
	XOR
	OR
)

func (op Operator) invoke(a bool, b bool) bool {
	switch op {
	case AND:
		return a && b
	case XOR:
		return a != b
	case OR:
		return a || b
	default:
		panic("invalid operator...")
	}
}

func fromString(str string) (Operator, error) {
	switch str {
	case "XOR":
		return XOR, nil
	case "OR":
		return OR, nil
	case "AND":
		return AND, nil
	default:
		return 0, fmt.Errorf("invalid operator '%s'", str)
	}
}

func (op Operator) toString() string {
	switch op {
	case AND:
		return "AND"
	case XOR:
		return "XOR"
	case OR:
		return "OR"
	default:
		panic("invalid operator...")
	}
}

type Wire struct {
	label string
	value *bool
	setBy *Gate
}

type Gate struct {
	inputA   *Wire
	inputB   *Wire
	operator Operator
	output   *Wire
}

type Board struct {
	wires map[string]*Wire
	gates []*Gate
	// connections map[string][]*Gate
}

func (b Board) getWireValue(label string) bool {
	wire := b.wires[label]
	if wire.value == nil {
		gate := wire.setBy
		inputA := b.getWireValue(gate.inputA.label)
		inputB := b.getWireValue(gate.inputB.label)
		value := gate.operator.invoke(inputA, inputB)
		wire.value = &value
	}
	return *wire.value
}

func parseInput(useRealInput bool) (*Board, error) {
	data, err := util.ReadInputMulti(24, useRealInput)
	if err != nil {
		return nil, err
	}
	if len(data) != 2 {
		return nil, fmt.Errorf("expected 2 parts of input")
	}

	// Read out input wires
	wires := make(map[string]*Wire)
	for _, line := range data[0] {
		split := strings.Split(line, ":")
		varName := split[0]
		varValue := strings.Contains(split[1], "1")
		wires[varName] = &Wire{varName, &varValue, nil}
	}

	// Helper that finds existing wire or creates new one if we didn't see it before
	findWireOrCreate := func(label string) *Wire {
		var wire *Wire
		wire, exists := wires[label]
		if !exists {
			wire = &Wire{label, nil, nil}
			wires[label] = wire
		}
		return wire
	}

	// Read out all the gates
	gates := make([]*Gate, 0)
	for _, line := range data[1] {
		split := strings.Split(line, " ")

		wireInA := findWireOrCreate(split[0])
		wireInB := findWireOrCreate(split[2])
		wireOut := findWireOrCreate(split[4])

		op, err := fromString(split[1])
		if err != nil {
			return nil, err
		}

		gate := Gate{wireInA, wireInB, op, wireOut}
		gates = append(gates, &gate)
		wireOut.setBy = &gate
	}

	return &Board{wires: wires, gates: gates}, nil
}

// =============== manually checking part 2 =======================

func SolvePart2(useRealInput bool) (int, error) {
	board, err := parseInput(useRealInput)
	if err != nil {
		return 0, err
	}

	// Idea is that we have a logic circuit representing the sum of two binary inputs.
	// The way that works is that we have so-called "full adders" for each input bits x and y and some input carry bit c_in, which then produces an output bit z and a carry bit c_out
	//
	// A full adder consists of 2 XOR's, 2 AND's and a OR:
	// x    XOR  y -> a1      (XOR_INPUT)
	// x    AND  y -> a2      (AND_INPUT)
	// c_in AND a1 -> a3      (AND_INTERNAL)
	// c_in XOR a1 -> z       (XOR_OUTPUT)
	// a2   OR  a3 -> c_out   (OR_CARRY)
	//
	// This is then repeated for all inputs, with the last most carry bit (for x44 and y44) being output z45
	//
	// Only exception is the least significant input bits, that is a "half adder":
	// x    XOR  y -> z
	// x    AND  y -> c_out

	// First visualize the graph with graphviz
	fmt.Print(board.toGraphString()) // output see: https://dreampuf.github.io/GraphvizOnline/?engine=dot#digraph%20G%20%7B%0A%0Asvb-%3EXOR_nwq%0Afsw-%3EXOR_nwq%0AXOR_nwq-%3Enwq%0Ay40-%3EXOR_hsh%0Ax40-%3EXOR_hsh%0AXOR_hsh-%3Ehsh%0Apjd-%3EXOR_z35%0Absk-%3EXOR_z35%0AXOR_z35-%3Ez35%0Atmt-%3EOR_qcv%0Adbj-%3EOR_qcv%0AOR_qcv-%3Eqcv%0Afvw-%3EAND_bms%0Andj-%3EAND_bms%0AAND_bms-%3Ebms%0Ay09-%3EXOR_cpt%0Ax09-%3EXOR_cpt%0AXOR_cpt-%3Ecpt%0Awcj-%3EXOR_z33%0Anct-%3EXOR_z33%0AXOR_z33-%3Ez33%0Ax20-%3EXOR_msm%0Ay20-%3EXOR_msm%0AXOR_msm-%3Emsm%0Athq-%3EAND_nfh%0Abmg-%3EAND_nfh%0AAND_nfh-%3Enfh%0Acjb-%3EOR_z18%0Akqr-%3EOR_z18%0AOR_z18-%3Ez18%0Ax01-%3EXOR_mtd%0Ay01-%3EXOR_mtd%0AXOR_mtd-%3Emtd%0Ay23-%3EAND_pcq%0Ax23-%3EAND_pcq%0AAND_pcq-%3Epcq%0Ay11-%3EAND_fvv%0Ax11-%3EAND_fvv%0AAND_fvv-%3Efvv%0Ay03-%3EAND_vmj%0Ax03-%3EAND_vmj%0AAND_vmj-%3Evmj%0Avsm-%3EOR_psj%0Anqs-%3EOR_psj%0AOR_psj-%3Epsj%0Apsj-%3EXOR_z10%0Arqp-%3EXOR_z10%0AXOR_z10-%3Ez10%0Ay06-%3EAND_gnt%0Ax06-%3EAND_gnt%0AAND_gnt-%3Egnt%0Ay13-%3EAND_jrk%0Ax13-%3EAND_jrk%0AAND_jrk-%3Ejrk%0Anhq-%3EXOR_z39%0Avjj-%3EXOR_z39%0AXOR_z39-%3Ez39%0Adqq-%3EOR_ntg%0Agkj-%3EOR_ntg%0AOR_ntg-%3Entg%0Ax28-%3EXOR_hjc%0Ay28-%3EXOR_hjc%0AXOR_hjc-%3Ehjc%0Abff-%3EAND_mvd%0Ahsh-%3EAND_mvd%0AAND_mvd-%3Emvd%0Ax18-%3EXOR_gdw%0Ay18-%3EXOR_gdw%0AXOR_gdw-%3Egdw%0Abqc-%3EXOR_z16%0Amdd-%3EXOR_z16%0AXOR_z16-%3Ez16%0Ay01-%3EAND_dsm%0Ax01-%3EAND_dsm%0AAND_dsm-%3Edsm%0Ay44-%3EAND_qmh%0Ax44-%3EAND_qmh%0AAND_qmh-%3Eqmh%0Acbs-%3EAND_sbd%0Apfr-%3EAND_sbd%0AAND_sbd-%3Esbd%0Ax39-%3EXOR_nhq%0Ay39-%3EXOR_nhq%0AXOR_nhq-%3Enhq%0Addf-%3EAND_dbj%0Apvc-%3EAND_dbj%0AAND_dbj-%3Edbj%0Ay37-%3EAND_shr%0Ax37-%3EAND_shr%0AAND_shr-%3Eshr%0Arpv-%3EAND_wjv%0Awpq-%3EAND_wjv%0AAND_wjv-%3Ewjv%0Adtt-%3EXOR_z17%0Aqgt-%3EXOR_z17%0AXOR_z17-%3Ez17%0Ay24-%3EXOR_jdw%0Ax24-%3EXOR_jdw%0AXOR_jdw-%3Ejdw%0Apvd-%3EOR_qgt%0Actn-%3EOR_qgt%0AOR_qgt-%3Eqgt%0Awcj-%3EAND_gkj%0Anct-%3EAND_gkj%0AAND_gkj-%3Egkj%0Ajvp-%3EAND_rrt%0Asmc-%3EAND_rrt%0AAND_rrt-%3Errt%0Ax29-%3EAND_fqs%0Ay29-%3EAND_fqs%0AAND_fqs-%3Efqs%0Anwg-%3EXOR_mdb%0Afsf-%3EXOR_mdb%0AXOR_mdb-%3Emdb%0Apcq-%3EOR_sjg%0Amqc-%3EOR_sjg%0AOR_sjg-%3Esjg%0Akjd-%3EOR_fsw%0Adwf-%3EOR_fsw%0AOR_fsw-%3Efsw%0Ajrk-%3EOR_mcw%0Arww-%3EOR_mcw%0AOR_mcw-%3Emcw%0Amkq-%3EOR_fgg%0Avmf-%3EOR_fgg%0AOR_fgg-%3Efgg%0Andj-%3EXOR_z19%0Afvw-%3EXOR_z19%0AXOR_z19-%3Ez19%0Afdq-%3EOR_phs%0Arrt-%3EOR_phs%0AOR_phs-%3Ephs%0Agsc-%3EOR_hrn%0Agnt-%3EOR_hrn%0AOR_hrn-%3Ehrn%0Ay08-%3EXOR_kvn%0Ax08-%3EXOR_kvn%0AXOR_kvn-%3Ekvn%0Arjn-%3EXOR_z15%0Asvf-%3EXOR_z15%0AXOR_z15-%3Ez15%0Ajqn-%3EOR_rpv%0Arwg-%3EOR_rpv%0AOR_rpv-%3Erpv%0Ax06-%3EXOR_rqd%0Ay06-%3EXOR_rqd%0AXOR_rqd-%3Erqd%0Anwg-%3EAND_z22%0Afsf-%3EAND_z22%0AAND_z22-%3Ez22%0Ay27-%3EXOR_knv%0Ax27-%3EXOR_knv%0AXOR_knv-%3Eknv%0Adnn-%3EXOR_z11%0Ahrd-%3EXOR_z11%0AXOR_z11-%3Ez11%0Ay42-%3EAND_nwp%0Ax42-%3EAND_nwp%0AAND_nwp-%3Enwp%0Avmj-%3EOR_bqv%0Abtd-%3EOR_bqv%0AOR_bqv-%3Ebqv%0Ay40-%3EAND_dtp%0Ax40-%3EAND_dtp%0AAND_dtp-%3Edtp%0Ay12-%3EXOR_skg%0Ax12-%3EXOR_skg%0AXOR_skg-%3Eskg%0Ax30-%3EAND_dmf%0Ay30-%3EAND_dmf%0AAND_dmf-%3Edmf%0Abmg-%3EXOR_z21%0Athq-%3EXOR_z21%0AXOR_z21-%3Ez21%0Ax25-%3EXOR_smc%0Ay25-%3EXOR_smc%0AXOR_smc-%3Esmc%0Asjg-%3EAND_nhj%0Ajdw-%3EAND_nhj%0AAND_nhj-%3Enhj%0Ax15-%3EAND_spm%0Ay15-%3EAND_spm%0AAND_spm-%3Espm%0Ay41-%3EAND_tmt%0Ax41-%3EAND_tmt%0AAND_tmt-%3Etmt%0Avmq-%3EXOR_z37%0Ahmw-%3EXOR_z37%0AXOR_z37-%3Ez37%0Adwq-%3EOR_cbs%0Atkd-%3EOR_cbs%0AOR_cbs-%3Ecbs%0Ajjg-%3EOR_hwd%0Afvv-%3EOR_hwd%0AOR_hwd-%3Ehwd%0Ax32-%3EXOR_npn%0Ay32-%3EXOR_npn%0AXOR_npn-%3Enpn%0Ajpq-%3EXOR_z02%0Awjg-%3EXOR_z02%0AXOR_z02-%3Ez02%0Arjb-%3EOR_gfc%0Akwm-%3EOR_gfc%0AOR_gfc-%3Egfc%0Ay31-%3EAND_wfh%0Ax31-%3EAND_wfh%0AAND_wfh-%3Ewfh%0Amdd-%3EAND_ctn%0Abqc-%3EAND_ctn%0AAND_ctn-%3Ectn%0Awpq-%3EXOR_z05%0Arpv-%3EXOR_z05%0AXOR_z05-%3Ez05%0Ax35-%3EXOR_bsk%0Ay35-%3EXOR_bsk%0AXOR_bsk-%3Ebsk%0Amtd-%3EAND_jrp%0Appj-%3EAND_jrp%0AAND_jrp-%3Ejrp%0Ahrn-%3EAND_vds%0Ackm-%3EAND_vds%0AAND_vds-%3Evds%0Ax07-%3EXOR_ckm%0Ay07-%3EXOR_ckm%0AXOR_ckm-%3Eckm%0Ax05-%3EXOR_grf%0Ay05-%3EXOR_grf%0AXOR_grf-%3Egrf%0Ax07-%3EAND_hgj%0Ay07-%3EAND_hgj%0AAND_hgj-%3Ehgj%0Awjv-%3EOR_chd%0Agrf-%3EOR_chd%0AOR_chd-%3Echd%0Ax20-%3EAND_prv%0Ay20-%3EAND_prv%0AAND_prv-%3Eprv%0Ajqg-%3EOR_tpj%0Adhg-%3EOR_tpj%0AOR_tpj-%3Etpj%0Apws-%3EOR_bff%0Abjg-%3EOR_bff%0AOR_bff-%3Ebff%0Ay24-%3EAND_rrr%0Ax24-%3EAND_rrr%0AAND_rrr-%3Errr%0Ay43-%3EXOR_vbm%0Ax43-%3EXOR_vbm%0AXOR_vbm-%3Evbm%0Ax03-%3EXOR_hsj%0Ay03-%3EXOR_hsj%0AXOR_hsj-%3Ehsj%0Arqd-%3EXOR_z06%0Achd-%3EXOR_z06%0AXOR_z06-%3Ez06%0Ajdw-%3EXOR_z24%0Asjg-%3EXOR_z24%0AXOR_z24-%3Ez24%0Affh-%3EXOR_fvw%0Agdw-%3EXOR_fvw%0AXOR_fvw-%3Efvw%0Agjb-%3EAND_jqj%0Adjh-%3EAND_jqj%0AAND_jqj-%3Ejqj%0Ahgq-%3EXOR_z30%0Afbm-%3EXOR_z30%0AXOR_z30-%3Ez30%0Ax21-%3EXOR_thq%0Ay21-%3EXOR_thq%0AXOR_thq-%3Ethq%0Agjb-%3EXOR_z38%0Adjh-%3EXOR_z38%0AXOR_z38-%3Ez38%0Armj-%3EOR_pjd%0Acbr-%3EOR_pjd%0AOR_pjd-%3Epjd%0Ay00-%3EXOR_z00%0Ax00-%3EXOR_z00%0AXOR_z00-%3Ez00%0Anpn-%3EAND_smh%0Avrp-%3EAND_smh%0AAND_smh-%3Esmh%0Ax32-%3EAND_whh%0Ay32-%3EAND_whh%0AAND_whh-%3Ewhh%0Aqcq-%3EOR_fmd%0Akjq-%3EOR_fmd%0AOR_fmd-%3Efmd%0Ax44-%3EXOR_pfr%0Ay44-%3EXOR_pfr%0AXOR_pfr-%3Epfr%0Absk-%3EAND_dwf%0Apjd-%3EAND_dwf%0AAND_dwf-%3Edwf%0Admf-%3EOR_jpb%0Agcw-%3EOR_jpb%0AOR_jpb-%3Ejpb%0Amtd-%3EXOR_z01%0Appj-%3EXOR_z01%0AXOR_z01-%3Ez01%0Asbs-%3EOR_dqr%0Amdb-%3EOR_dqr%0AOR_dqr-%3Edqr%0Ay13-%3EXOR_rnw%0Ax13-%3EXOR_rnw%0AXOR_rnw-%3Ernw%0Ahgj-%3EOR_wjp%0Avds-%3EOR_wjp%0AOR_wjp-%3Ewjp%0Avdd-%3EOR_vrp%0Awfh-%3EOR_vrp%0AOR_vrp-%3Evrp%0Ax22-%3EXOR_fsf%0Ay22-%3EXOR_fsf%0AXOR_fsf-%3Efsf%0Ax31-%3EXOR_qjc%0Ay31-%3EXOR_qjc%0AXOR_qjc-%3Eqjc%0Ay22-%3EAND_sbs%0Ax22-%3EAND_sbs%0AAND_sbs-%3Esbs%0Ay36-%3EAND_z36%0Ax36-%3EAND_z36%0AAND_z36-%3Ez36%0Abqv-%3EAND_rwg%0Afdk-%3EAND_rwg%0AAND_rwg-%3Erwg%0Akvm-%3EAND_krc%0Aqcv-%3EAND_krc%0AAND_krc-%3Ekrc%0Ahwk-%3EOR_dnn%0Abqk-%3EOR_dnn%0AOR_dnn-%3Ednn%0Askg-%3EAND_fbd%0Ahwd-%3EAND_fbd%0AAND_fbd-%3Efbd%0Aftr-%3EAND_fkw%0Amcw-%3EAND_fkw%0AAND_fkw-%3Efkw%0Abcc-%3EXOR_z13%0Arnw-%3EXOR_z13%0AXOR_z13-%3Ez13%0Ay16-%3EAND_pvd%0Ax16-%3EAND_pvd%0AAND_pvd-%3Epvd%0Aqcv-%3EXOR_z42%0Akvm-%3EXOR_z42%0AXOR_z42-%3Ez42%0Amhf-%3EAND_dhg%0Aphs-%3EAND_dhg%0AAND_dhg-%3Edhg%0Atpj-%3EAND_kjq%0Aknv-%3EAND_kjq%0AAND_kjq-%3Ekjq%0Ay00-%3EAND_ppj%0Ax00-%3EAND_ppj%0AAND_ppj-%3Eppj%0Adcn-%3EAND_wpr%0Afgg-%3EAND_wpr%0AAND_wpr-%3Ewpr%0Artt-%3EOR_nwg%0Anfh-%3EOR_nwg%0AOR_nwg-%3Enwg%0Ashr-%3EOR_gjb%0Awmg-%3EOR_gjb%0AOR_gjb-%3Egjb%0Ax23-%3EXOR_bnv%0Ay23-%3EXOR_bnv%0AXOR_bnv-%3Ebnv%0Ay04-%3EAND_jqn%0Ax04-%3EAND_jqn%0AAND_jqn-%3Ejqn%0Ay02-%3EXOR_jpq%0Ax02-%3EXOR_jpq%0AXOR_jpq-%3Ejpq%0Ay17-%3EAND_vvt%0Ax17-%3EAND_vvt%0AAND_vvt-%3Evvt%0Ay27-%3EAND_qcq%0Ax27-%3EAND_qcq%0AAND_qcq-%3Eqcq%0Ax34-%3EXOR_htp%0Ay34-%3EXOR_htp%0AXOR_htp-%3Ehtp%0Anwq-%3EOR_vmq%0Atvh-%3EOR_vmq%0AOR_vmq-%3Evmq%0Aqvr-%3EOR_vjj%0Ajqj-%3EOR_vjj%0AOR_vjj-%3Evjj%0Ax17-%3EXOR_dtt%0Ay17-%3EXOR_dtt%0AXOR_dtt-%3Edtt%0Ahsj-%3EXOR_z03%0Agfc-%3EXOR_z03%0AXOR_z03-%3Ez03%0Ay16-%3EXOR_bqc%0Ax16-%3EXOR_bqc%0AXOR_bqc-%3Ebqc%0Anrr-%3EAND_vvb%0Amsm-%3EAND_vvb%0AAND_vvb-%3Evvb%0Anhj-%3EOR_jvp%0Arrr-%3EOR_jvp%0AOR_jvp-%3Ejvp%0Amvd-%3EOR_pvc%0Adtp-%3EOR_pvc%0AOR_pvc-%3Epvc%0Avmq-%3EAND_wmg%0Ahmw-%3EAND_wmg%0AAND_wmg-%3Ewmg%0Akvn-%3EXOR_z08%0Awjp-%3EXOR_z08%0AXOR_z08-%3Ez08%0Ax21-%3EAND_rtt%0Ay21-%3EAND_rtt%0AAND_rtt-%3Ertt%0Aqgt-%3EAND_bct%0Adtt-%3EAND_bct%0AAND_bct-%3Ebct%0Ax19-%3EAND_tdg%0Ay19-%3EAND_tdg%0AAND_tdg-%3Etdg%0Ax41-%3EXOR_ddf%0Ay41-%3EXOR_ddf%0AXOR_ddf-%3Eddf%0Atvb-%3EOR_bfv%0Agbn-%3EOR_bfv%0AOR_bfv-%3Ebfv%0Apvc-%3EXOR_z41%0Addf-%3EXOR_z41%0AXOR_z41-%3Ez41%0Ax05-%3EAND_wpq%0Ay05-%3EAND_wpq%0AAND_wpq-%3Ewpq%0Aftr-%3EXOR_z14%0Amcw-%3EXOR_z14%0AXOR_z14-%3Ez14%0Ay19-%3EXOR_ndj%0Ax19-%3EXOR_ndj%0AXOR_ndj-%3Endj%0Ax26-%3EXOR_mhf%0Ay26-%3EXOR_mhf%0AXOR_mhf-%3Emhf%0Abff-%3EXOR_z40%0Ahsh-%3EXOR_z40%0AXOR_z40-%3Ez40%0Ax43-%3EAND_tkd%0Ay43-%3EAND_tkd%0AAND_tkd-%3Etkd%0Afsw-%3EAND_tvh%0Asvb-%3EAND_tvh%0AAND_tvh-%3Etvh%0Atdg-%3EOR_nrr%0Abms-%3EOR_nrr%0AOR_nrr-%3Enrr%0Antg-%3EXOR_z34%0Ahtp-%3EXOR_z34%0AXOR_z34-%3Ez34%0Awpr-%3EOR_hgq%0Afqs-%3EOR_hgq%0AOR_hgq-%3Ehgq%0Avvt-%3EOR_ffh%0Abct-%3EOR_ffh%0AOR_ffh-%3Effh%0Ay25-%3EAND_fdq%0Ax25-%3EAND_fdq%0AAND_fdq-%3Efdq%0Adcn-%3EXOR_z29%0Afgg-%3EXOR_z29%0AXOR_z29-%3Ez29%0Ahtp-%3EAND_cbr%0Antg-%3EAND_cbr%0AAND_cbr-%3Ecbr%0Abnv-%3EAND_mqc%0Adqr-%3EAND_mqc%0AAND_mqc-%3Emqc%0Ay04-%3EXOR_fdk%0Ax04-%3EXOR_fdk%0AXOR_fdk-%3Efdk%0Aprv-%3EOR_bmg%0Avvb-%3EOR_bmg%0AOR_bmg-%3Ebmg%0Ay33-%3EAND_dqq%0Ax33-%3EAND_dqq%0AAND_dqq-%3Edqq%0Ay18-%3EAND_cjb%0Ax18-%3EAND_cjb%0AAND_cjb-%3Ecjb%0Aqjc-%3EXOR_z31%0Ajpb-%3EXOR_z31%0AXOR_z31-%3Ez31%0Ajpb-%3EAND_vdd%0Aqjc-%3EAND_vdd%0AAND_vdd-%3Evdd%0Afmd-%3EAND_mkq%0Ahjc-%3EAND_mkq%0AAND_mkq-%3Emkq%0Afbd-%3EOR_bcc%0Arvn-%3EOR_bcc%0AOR_bcc-%3Ebcc%0Acbs-%3EXOR_z44%0Apfr-%3EXOR_z44%0AXOR_z44-%3Ez44%0Ax34-%3EAND_rmj%0Ay34-%3EAND_rmj%0AAND_rmj-%3Ermj%0Ax30-%3EXOR_fbm%0Ay30-%3EXOR_fbm%0AXOR_fbm-%3Efbm%0Ahrn-%3EXOR_z07%0Ackm-%3EXOR_z07%0AXOR_z07-%3Ez07%0Adnn-%3EAND_jjg%0Ahrd-%3EAND_jjg%0AAND_jjg-%3Ejjg%0Anqg-%3EXOR_z43%0Avbm-%3EXOR_z43%0AXOR_z43-%3Ez43%0Agpc-%3EOR_rjn%0Afkw-%3EOR_rjn%0AOR_rjn-%3Erjn%0Ay37-%3EXOR_hmw%0Ax37-%3EXOR_hmw%0AXOR_hmw-%3Ehmw%0Ahgq-%3EAND_gcw%0Afbm-%3EAND_gcw%0AAND_gcw-%3Egcw%0Ay02-%3EAND_rjb%0Ax02-%3EAND_rjb%0AAND_rjb-%3Erjb%0Ax10-%3EAND_hwk%0Ay10-%3EAND_hwk%0AAND_hwk-%3Ehwk%0Ay38-%3EAND_qvr%0Ax38-%3EAND_qvr%0AAND_qvr-%3Eqvr%0Ax14-%3EXOR_ftr%0Ay14-%3EXOR_ftr%0AXOR_ftr-%3Eftr%0Avjj-%3EAND_pws%0Anhq-%3EAND_pws%0AAND_pws-%3Epws%0Anwp-%3EOR_nqg%0Akrc-%3EOR_nqg%0AOR_nqg-%3Enqg%0Ax11-%3EXOR_hrd%0Ay11-%3EXOR_hrd%0AXOR_hrd-%3Ehrd%0Abfv-%3EXOR_z09%0Acpt-%3EXOR_z09%0AXOR_z09-%3Ez09%0Agdw-%3EAND_kqr%0Affh-%3EAND_kqr%0AAND_kqr-%3Ekqr%0Ay26-%3EAND_jqg%0Ax26-%3EAND_jqg%0AAND_jqg-%3Ejqg%0Ax15-%3EXOR_svf%0Ay15-%3EXOR_svf%0AXOR_svf-%3Esvf%0Ax33-%3EXOR_nct%0Ay33-%3EXOR_nct%0AXOR_nct-%3Enct%0Achd-%3EAND_gsc%0Arqd-%3EAND_gsc%0AAND_gsc-%3Egsc%0Apfb-%3EOR_mdd%0Aspm-%3EOR_mdd%0AOR_mdd-%3Emdd%0Anpn-%3EXOR_z32%0Avrp-%3EXOR_z32%0AXOR_z32-%3Ez32%0Arjn-%3EAND_pfb%0Asvf-%3EAND_pfb%0AAND_pfb-%3Epfb%0Avbm-%3EAND_dwq%0Anqg-%3EAND_dwq%0AAND_dwq-%3Edwq%0Ax12-%3EAND_rvn%0Ay12-%3EAND_rvn%0AAND_rvn-%3Ervn%0Ax29-%3EXOR_dcn%0Ay29-%3EXOR_dcn%0AXOR_dcn-%3Edcn%0Ax38-%3EXOR_djh%0Ay38-%3EXOR_djh%0AXOR_djh-%3Edjh%0Askg-%3EXOR_z12%0Ahwd-%3EXOR_z12%0AXOR_z12-%3Ez12%0Ay14-%3EAND_gpc%0Ax14-%3EAND_gpc%0AAND_gpc-%3Egpc%0Aqmh-%3EOR_z45%0Asbd-%3EOR_z45%0AOR_z45-%3Ez45%0Akvn-%3EAND_tvb%0Awjp-%3EAND_tvb%0AAND_tvb-%3Etvb%0Ajvp-%3EXOR_z25%0Asmc-%3EXOR_z25%0AXOR_z25-%3Ez25%0Afdk-%3EXOR_z04%0Abqv-%3EXOR_z04%0AXOR_z04-%3Ez04%0Abfv-%3EAND_nqs%0Acpt-%3EAND_nqs%0AAND_nqs-%3Enqs%0Ay09-%3EAND_vsm%0Ax09-%3EAND_vsm%0AAND_vsm-%3Evsm%0Ay10-%3EXOR_rqp%0Ax10-%3EXOR_rqp%0AXOR_rqp-%3Erqp%0Asmh-%3EOR_wcj%0Awhh-%3EOR_wcj%0AOR_wcj-%3Ewcj%0Ax39-%3EAND_bjg%0Ay39-%3EAND_bjg%0AAND_bjg-%3Ebjg%0Ay08-%3EAND_gbn%0Ax08-%3EAND_gbn%0AAND_gbn-%3Egbn%0Ax36-%3EXOR_svb%0Ay36-%3EXOR_svb%0AXOR_svb-%3Esvb%0Apsj-%3EAND_bqk%0Arqp-%3EAND_bqk%0AAND_bqk-%3Ebqk%0Ay35-%3EAND_kjd%0Ax35-%3EAND_kjd%0AAND_kjd-%3Ekjd%0Amhf-%3EXOR_z26%0Aphs-%3EXOR_z26%0AXOR_z26-%3Ez26%0Ay28-%3EAND_vmf%0Ax28-%3EAND_vmf%0AAND_vmf-%3Evmf%0Ay42-%3EXOR_kvm%0Ax42-%3EXOR_kvm%0AXOR_kvm-%3Ekvm%0Abnv-%3EXOR_z23%0Adqr-%3EXOR_z23%0AXOR_z23-%3Ez23%0Agfc-%3EAND_btd%0Ahsj-%3EAND_btd%0AAND_btd-%3Ebtd%0Anrr-%3EXOR_z20%0Amsm-%3EXOR_z20%0AXOR_z20-%3Ez20%0Adsm-%3EOR_wjg%0Ajrp-%3EOR_wjg%0AOR_wjg-%3Ewjg%0Arnw-%3EAND_rww%0Abcc-%3EAND_rww%0AAND_rww-%3Erww%0Awjg-%3EAND_kwm%0Ajpq-%3EAND_kwm%0AAND_kwm-%3Ekwm%0Afmd-%3EXOR_z28%0Ahjc-%3EXOR_z28%0AXOR_z28-%3Ez28%0Atpj-%3EXOR_z27%0Aknv-%3EXOR_z27%0AXOR_z27-%3Ez27%0A%7D

	// Now, we are going to do some consistency checks:
	invalidGates := make([]*Gate, 0)
	gateRoles := make(map[*Gate]GateRole)
	for _, gate := range board.gates {
		switch {
		// half adder for first bit
		case gate.operator == XOR && (gate.inputsAre("x00", "y00") && gate.output.label == "z00"):
			gateRoles[gate] = XOR_INPUT_0
		case gate.operator == AND && (gate.inputsAre("x00", "y00") && gate.output.label[0] != 'z'):
			gateRoles[gate] = AND_INPUT_0
		// full adder for other bits
		case gate.operator == XOR && (gate.inputsStartWith('x', 'y') && gate.output.label[0] != 'z'):
			gateRoles[gate] = XOR_INPUT_N
		case gate.operator == XOR && (gate.inputsDontStartWith('x', 'y') && gate.output.label[0] == 'z'):
			gateRoles[gate] = XOR_OUTPUT_N
		case gate.operator == AND && (gate.inputsStartWith('x', 'y') && gate.output.label[0] != 'z'):
			gateRoles[gate] = AND_INPUT_N
		case gate.operator == AND && (gate.inputsDontStartWith('x', 'y') && gate.output.label[0] != 'z'):
			gateRoles[gate] = AND_INTERNAL_N
		case gate.operator == OR && (gate.inputsDontStartWith('x', 'y') && gate.output.label[0] != 'z'):
			gateRoles[gate] = OR_CARRY_N
		// final output bit should be the carry of the last addition
		case gate.operator == OR && (gate.inputsDontStartWith('x', 'y') && gate.output.label == "z45"):
			gateRoles[gate] = OR_CARRY_FINAL
		default:
			gateRoles[gate] = INVALID
			invalidGates = append(invalidGates, gate)
		}
	}
	// At this point there is 3 pairs of gates + output wires that are invalid
	// Manual inspection shows we need to do these swaps:
	// nwq <-> z36
	// fvw <-> z18
	// mdb <-> z22

	// There is one more swap we have to do.
	// To find candidates we search for all XOR "input" gates (XOR's that connect directly to x's and y's) that don't connect to an XOR "output" gate (XOR's that connect directly to z's)
	for gate, gateRole := range gateRoles {
		if gateRole == XOR_INPUT_N {
			outLabel := gate.output.label

			outputXorCounter := 0
			for otherGate, otherRole := range gateRoles {
				if otherGate.inputA.label == outLabel || otherGate.inputB.label == outLabel {
					if otherRole == XOR_OUTPUT_N {
						outputXorCounter += 1
					}
				}
			}

			if outputXorCounter != 1 {
				invalidGates = append(invalidGates, gate)
			}
		}
	}
	for _, gate := range invalidGates {
		fmt.Println(gate.output.label)
	}
	// This has found 4 more invalid configurations:
	// grf
	// fsf
	// gdw
	// svb
	// Manual inspection shows that fsf, gdw and svb are actually fixed already if we do above found swaps
	// But grf is a new one! If we just check the graph we see that it should be swapped with wpq, so we add it to our list:
	// grf <-> wpq

	// So now we have our answer: fvw,grf,mdb,nwq,wpq,z18,z22,z36
	return 0, nil
}

// Print the board as a string that can be put into graphviz
func (b Board) toGraphString() string {
	var sb strings.Builder

	for _, gate := range b.gates {
		gateName := fmt.Sprintf("%s_%s", gate.operator.toString(), gate.output.label)
		sb.WriteString(fmt.Sprintf("%s->%s\n", gate.inputA.label, gateName))
		sb.WriteString(fmt.Sprintf("%s->%s\n", gate.inputB.label, gateName))
		sb.WriteString(fmt.Sprintf("%s->%s\n", gateName, gate.output.label))
	}
	return sb.String()
}

type GateRole int

const (
	XOR_INPUT_0 GateRole = iota
	XOR_INPUT_N
	XOR_OUTPUT_N
	AND_INPUT_0
	AND_INPUT_N
	AND_INTERNAL_N
	OR_CARRY_N
	OR_CARRY_FINAL
	INVALID
)

func (gate *Gate) inputsStartWith(a, b byte) bool {
	return (gate.inputA.label[0] == a && gate.inputB.label[0] == b) || (gate.inputA.label[0] == b && gate.inputB.label[0] == a)
}

func (gate *Gate) inputsDontStartWith(a, b byte) bool {
	return gate.inputA.label[0] != a && gate.inputB.label[0] != b && gate.inputA.label[0] != b && gate.inputB.label[0] != a
}

func (gate *Gate) inputsAre(a, b string) bool {
	return (gate.inputA.label == a && gate.inputB.label == b) || (gate.inputA.label == b && gate.inputB.label == a)
}
