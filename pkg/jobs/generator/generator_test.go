package generator_test

import (
	"github.com/bellh14/DFRDesignManager/pkg/jobs/generator"
	"github.com/bellh14/DFRDesignManager/pkg/types"
	"testing"
)

func TestGenerateJobScript(t *testing.T) {
	jobScriptInputs := types.JobSubmissionType{
		WorkingDir: ".",
		Ntasks:     4,
		Path:       "/opt/Siemens/17.04.008-R8/STAR-CCM+17.04.008-R8/star/bin/",
		PodKey:     "1234-5678-9012-3456",
		JavaMacro:  "macro.java",
		SimFile:    "sim.sim",
		JobNumber:  1,
	}
	generator.GenerateJobScript(jobScriptInputs)

}
