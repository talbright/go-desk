package convey

import (
	"fmt"
	"strings"

	"github.com/smartystreets/goconvey/convey/reporting"
)

type scope struct {
	name       string
	title      string
	action     *action
	children   map[string]*scope
	birthOrder []*scope
	child      int
	resets     map[string]*action
	panicked   bool
	reporter   reporting.Reporter
	report     *reporting.ScopeReport
}

func (parent *scope) adopt(child *scope) {
	i := parent.getChildIndex(child)

	if i == -1 {
		parent.children[child.name] = child
		parent.birthOrder = append(parent.birthOrder, child)
	} else {
		/* We need to replace the action to retain the closed over variables from
		   the specific invocation of the parent scope, enabling the enclosing
		   parent scope to serve as a set-up for the child scope */
		parent.birthOrder[i].action = child.action
	}
}

func (parent *scope) getChildIndex(child *scope) int {
	for i, ordered := range parent.birthOrder {
		if ordered.name == child.name && ordered.title == child.title {
			return i
		}
	}

	return -1
}

func (self *scope) registerReset(action *action) {
	self.resets[action.name] = action
}

func (self *scope) visited() bool {
	return self.panicked || self.child >= len(self.birthOrder)
}

func (parent *scope) visit(runner *runner) {
	runner.active = parent
	defer parent.exit()

	parent.enter()
	parent.action.Invoke()
	parent.visitChildren(runner)
}
func (parent *scope) enter() {
	parent.reporter.Enter(parent.report)
}
func (parent *scope) visitChildren(runner *runner) {
	if len(parent.birthOrder) == 0 {
		parent.cleanup()
	} else {
		parent.visitChild(runner)
	}
}
func (parent *scope) visitChild(runner *runner) {
	child := parent.birthOrder[parent.child]
	child.visit(runner)

	parent.cleanup()

	if child.visited() {
		parent.child++
	}
}
func (parent *scope) cleanup() {
	for _, reset := range parent.resets {
		reset.Invoke()
	}
}
func (parent *scope) exit() {
	if problem := recover(); problem != nil {
		if strings.HasPrefix(fmt.Sprintf("%v", problem), extraGoTest) {
			panic(problem)
		}
		if problem != failureHalt {
			parent.reporter.Report(reporting.NewErrorReport(problem))
		}
		parent.panicked = true
	}
	parent.reporter.Exit()
}

func newScope(entry *registration, reporter reporting.Reporter) *scope {
	self := new(scope)
	self.reporter = reporter
	self.name = entry.action.name
	self.title = entry.Situation
	self.action = entry.action
	self.children = make(map[string]*scope)
	self.birthOrder = []*scope{}
	self.resets = make(map[string]*action)
	self.report = reporting.NewScopeReport(self.title, self.name)
	return self
}
