package file

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Processor struct {
	inpath  string
	outpath string
}

// New creates and initializes a new Processor for handling input and output file paths.
// It ensures that the input and output paths are not the same.
// Returns a Processor to manage file transformations and an error if the paths are identical.
func New(inpath, outpath string) (*Processor, error) {
	if inpath == outpath {
		return nil, fmt.Errorf("input and output paths must be different")
	}

	return &Processor{
		inpath:  inpath,
		outpath: outpath,
	}, nil
}

// Process reads a YAML file, deserializes its contents, and writes them in YAML format to an output file.
// It manages input and output file paths defined in the Processor, ensuring proper encoding and error handling.
// Returns an error if file operations fail, YAML parsing is unsuccessful, or encoding to the output file encounters an issue.
func (p *Processor) Process() error {
	data, err := os.ReadFile(p.inpath)
	if err != nil {
		return fmt.Errorf("failed to read input file %s: %w", p.inpath, err)
	}

	var node yaml.Node
	if err := yaml.Unmarshal(data, &node); err != nil {
		return fmt.Errorf("failed to parse YAML from input file %s: %w", p.inpath, err)
	}

	outFile, err := os.OpenFile(p.outpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o644)
	if err != nil {
		return fmt.Errorf("failed to open output file %s: %w", p.outpath, err)
	}
	defer outFile.Close()

	encoder := yaml.NewEncoder(outFile)

	if err := encoder.Encode(&node); err != nil {
		return fmt.Errorf("failed to encode YAML to output file %s: %w", p.outpath, err)
	}

	return nil
}
