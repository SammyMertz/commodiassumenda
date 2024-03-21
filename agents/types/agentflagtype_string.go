// Code generated by "stringer -type=AgentFlagType -linecomment"; DO NOT EDIT.

package types

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[AgentFlagUnknown-0]
	_ = x[AgentFlagActive-1]
	_ = x[AgentFlagUnstaking-2]
	_ = x[AgentFlagResting-3]
	_ = x[AgentFlagFraudulent-4]
	_ = x[AgentFlagSlashed-5]
}

const _AgentFlagType_name = "AgentFlagUnknownAgentFlagActiveAgentFlagUnstakingAgentFlagRestingAgentFlagFraudulentAgentFlagSlashed"

var _AgentFlagType_index = [...]uint8{0, 16, 31, 49, 65, 84, 100}

func (i AgentFlagType) String() string {
	if i >= AgentFlagType(len(_AgentFlagType_index)-1) {
		return "AgentFlagType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _AgentFlagType_name[_AgentFlagType_index[i]:_AgentFlagType_index[i+1]]
}
