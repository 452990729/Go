package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
)

var (
	h   bool
	in  string
	out string
)

func init() {
	flag.BoolVar(&h, "h", false, "this help")
	flag.StringVar(&in, "in", "./", "file path or dir path, if it is dir, recursive the whole dir. set `input`.")
	flag.StringVar(&out, "out", "", "if you want to remain the ori file, you need to define the output path. default do not remain the ori file and replace it. set `outpath`.")
	//	flag.Usage = usage
}

func makeXlsx(file, out string) {
	_, err := excelize.OpenFile(file)
	if err != nil {
		dir, fl := path.Split(file)
		lb := strings.Split(fl, ".")
		lb = lb[:len(lb)-1]
		if out != "" {
			dir = out
		}
		newFl := path.Join(dir, strings.Join(lb, ".")+".tmp.xlsx")
		newFlR := path.Join(dir, strings.Join(lb, ".")+".xlsx")
		fn, err := os.Open(file)
		if err != nil {
			fmt.Println(file, " can not read")
		} else {
			bf := bufio.NewScanner(fn)
			buf := make([]byte, 1000000)
			bf.Buffer(buf, bufio.MaxScanTokenSize)
			var cell = [...]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "AA", "AB", "AC", "AD", "AE", "AF", "AG", "AH", "AI", "AJ", "AK", "AL", "AM", "AN", "AO", "AP", "AQ", "AR", "AS", "AT", "AU", "AV", "AW", "AX", "AY", "AZ", "BA", "BB", "BC", "BD", "BE", "BF", "BG", "BH", "BI", "BJ", "BK", "BL", "BM", "BN", "BO", "BP", "BQ", "BR", "BS", "BT", "BU", "BV", "BW", "BX", "BY", "BZ", "CA", "CB", "CC", "CD", "CE", "CF", "CG", "CH", "CI", "CJ", "CK", "CL", "CM", "CN", "CO", "CP", "CQ", "CR", "CS", "CT", "CU", "CV", "CW", "CX", "CY", "CZ", "DA", "DB", "DC", "DD", "DE", "DF", "DG", "DH", "DI", "DJ", "DK", "DL", "DM", "DN", "DO", "DP", "DQ", "DR", "DS", "DT", "DU", "DV", "DW", "DX", "DY", "DZ"}
			m := 0
			xlsx := excelize.NewFile()
			log.Println("Processing", file)
			for bf.Scan() {
				m++
				slice_line := strings.Split(bf.Text(), "\t")
				for index, value := range slice_line {
					xlsx.SetCellValue("Sheet1", cell[index]+strconv.Itoa(m), value)
				}
			}
			err := bf.Err()
			if err != nil {
				log.Fatal(err)
			}
			if out == "" {
				os.Remove(file)
			}
			err = xlsx.SaveAs(newFl)
			if err != nil {
				fmt.Println(err)
			}
			os.Rename(newFl, newFlR)
		}
		defer func() {
			err := fn.Close()
			if err != nil {
				log.Fatal(err)
			}
		}()

	}
}

func getXls(inpath string, files []string) []string {
	cl, _ := ioutil.ReadDir(inpath)
	for _, f := range cl {
		if f.IsDir() {
			cll := path.Join(inpath, f.Name())
			files = getXls(cll, files)
		} else {
			if strings.HasSuffix(f.Name(), ".xls") || strings.HasSuffix(f.Name(), ".xlsx") {
				files = append(files, path.Join(inpath, f.Name()))
			}
		}
	}
	return files
}

func main() {
	flag.Parse()
	if h {
		fmt.Println("This program is used to convert false xls/xlsx suffix files to real xlsx files, Up to 104 columns of file are supported.\nThe input can be file or dir, if it is dir, then it will recursive the whole dir.\nIt replace the ori file default, but you can define the outpath to save the outfile and remain the ori file.")
		flag.Usage()
	}
	fl, err := os.Stat(in)
	if err != nil {
		fmt.Println("file or dir not exists")
		return
	}
	if fl.IsDir() {
		var files []string
		files = getXls(in, files)
		for i := 0; i < len(files); i++ {
			makeXlsx(files[i], out)
		}
	} else {
		makeXlsx(in, out)
	}
}
