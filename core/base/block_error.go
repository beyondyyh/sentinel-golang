package base

import "fmt"

// BlockError indicates the request was blocked by Sentinel.
type BlockError struct {
	blockType BlockType
	blockMsg  string

	rule          SentinelRule
	snapshotValue interface{}
}

func (e *BlockError) BlockMsg() string {
	return e.blockMsg
}

func (e *BlockError) BlockType() BlockType {
	return e.blockType
}

func (e *BlockError) TriggeredRule() SentinelRule {
	return e.rule
}

func (e *BlockError) TriggeredValue() interface{} {
	return e.snapshotValue
}

func NewBlockErrorFromDeepCopy(from *BlockError) *BlockError {
	return &BlockError{
		blockType:     from.blockType,
		blockMsg:      from.blockMsg,
		rule:          from.rule,
		snapshotValue: from.snapshotValue,
	}
}

func NewBlockError(blockType BlockType, blockMsg string) *BlockError {
	return &BlockError{blockType: blockType, blockMsg: blockMsg}
}

func NewBlockErrorWithCause(blockType BlockType, blockMsg string, rule SentinelRule, snapshot interface{}) *BlockError {
	return &BlockError{blockType: blockType, blockMsg: blockMsg, rule: rule, snapshotValue: snapshot}
}

func (e *BlockError) Error() string {
	return fmt.Sprintf("Sentinel block error: blockType:%s, message: %s, rule: %+v, snapshotValue: %+v",
		e.blockType, e.blockMsg, e.rule, e.snapshotValue)
}
