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

func New(inpath, outpath string) (*Processor, error) {
	if inpath == outpath {
		return nil, fmt.Errorf("input and output paths must be different")
	}

	return &Processor{
		inpath:  inpath,
		outpath: outpath,
	}, nil
}

// Process is a placeholder for the actual processing logic.
// It currently returns an error indicating that the method is not implemented.
func (p *Processor) Process() error {
	inFile, err := os.OpenFile(p.inpath, os.O_RDONLY, 0o644)
	if err != nil {
		return fmt.Errorf("failed to open input file %s: %w", p.inpath, err)
	}

	defer inFile.Close()

	var data any

	if err := yaml.NewDecoder(inFile).Decode(&data); err != nil {
		return fmt.Errorf("failed to decode input file %s: %w", p.inpath, err)
	}

	outFile, err := os.OpenFile(p.outpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o644)
	if err != nil {
		return fmt.Errorf("failed to open output file %s: %w", p.outpath, err)
	}

	defer outFile.Close()

	if err := yaml.NewEncoder(outFile).Encode(data); err != nil {
		return fmt.Errorf("failed to encode output file %s: %w", p.outpath, err)
	}

	return nil
}
