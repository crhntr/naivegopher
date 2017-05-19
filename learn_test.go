package naivegopher

import (
	"strings"
	"testing"
)

func TestLearn00(t *testing.T) {
	c := NewClassifier()
	c.Learn("bad", strings.NewReader("slow dirty old dull"))
	c.Learn("evil", strings.NewReader("horrid mean evil raceist trump"))
	c.Learn("good", strings.NewReader("fast clean new sharp education"))

	t.Log(c.CategoryNames())
	t.Log(c.ProbableCategoreies(strings.NewReader("home is where the heart is")))
	t.Log(c.ProbableCategoreies(strings.NewReader("the fast one")))
	t.Log(c.ProbableCategoreies(strings.NewReader("the slow one")))
	t.Log(c.ProbableCategoreies(strings.NewReader("the clean one")))
	t.Log(c.ProbableCategoreies(strings.NewReader("horrid mean evil raceist slow dirty old dull")))
}

func TestLearn01(t *testing.T) {
	panitumumab := []string{
		"Panitumumab in Metastatic Colorectal Cancer: The Importance of Tumour RAS Status",
		"Tumor penetration and epidermal growth factor receptor saturation by panitumumab correlate with antitumor activity in a preclinical model of human cancer",
		"Epidermal Growth Factor Receptor–Targeted Radioimmunotherapy of Human Head and Neck Cancer Xenografts Using 90Y-Labeled Fully Human Antibody Panitumumab",
		"In Vitro and In Vivo Analysis of Indocyanine Green-Labeled Panitumumab for Optical Imaging—A Cautionary Tale",
		"Randomized Phase Ib/II Trial of Rilotumumab or Ganitumab with Panitumumab versus Panitumumab Alone in Patients with Wild-type KRAS Metastatic Colorectal Cancer",
		"Immunogenicity of panitumumab in combination chemotherapy clinical trials",
		"Zirconium-89 Labeled Panitumumab: a Potential Immuno-PET Probe for HER1-Expressing Carcinomas",
		"Panitumumab Use in Metastatic Colorectal Cancer and Patterns of KRAS Testing: Results from a Europe-Wide Physician Survey and Medical Records Review",
		"Efficacy of CDK4 inhibition against sarcomas depends on their levels of CDK4 and p16ink4 mRNA",
		"Functional Dissection of the Epidermal Growth Factor Receptor Epitopes Targeted by Panitumumab and Cetuximab",
		"Panitumumab in Japanese Patients with Unresectable Colorectal Cancer: A Post-marketing Surveillance Study of 3085 Patients", "Preparation of clinical-grade 89Zr-panitumumab as a positron emission tomography biomarker for evaluating epidermal growth factor receptor-targeted therapy",
		"Development and Characterization of 89Zr-Labeled Panitumumab for Immuno–Positron Emission Tomographic Imaging of the Epidermal Growth Factor Receptor",
		"n analysis of the treatment effect of panitumumab on overall survival from a phase 3, randomized, controlled, multicenter trial (20020408) in patients with chemotherapy refractory metastatic colorectal cancer",
		"Epidermal growth factor receptor mutation mediates cross-resistance to panitumumab and cetuximab in gastrointestinal cancer",
		"Genomic markers of panitumumab resistance including ERBB2/HER2 in a phase II study of KRAS wild-type (wt) metastatic colorectal cancer (mCRC)",
	}

	palbociclib := []string{
		"Palbociclib inhibits epithelial-mesenchymal transition and metastasis in breast cancer via c-Jun/COX-2 signaling pathway",
		"Efflux Transporters at the Blood-Brain Barrier Limit Delivery and Efficacy of Cyclin-Dependent Kinase 4/6 Inhibitor Palbociclib (PD-0332991) in an Orthotopic Brain Tumor Model",
		"Inhibition of herpes simplex virus type 1 by the CDK6 inhibitor PD-0332991 (palbociclib) through the control of SAMHD1",
		"Mitigation of acute kidney injury by cell-cycle inhibitors that suppress both CDK4/6 and OCT2 functions",
		"Palbociclib for the Treatment of Estrogen Receptor–Positive, HER2-Negative Metastatic Breast Cancer",
		"Chemoproteomics reveals novel protein and lipid kinase targets of clinical CDK4/6 inhibitors in lung cancer",
		"Efficacy of SERD/SERM Hybrid-CDK4/6 inhibitor combinations in models of endocrine therapy resistant breast cancer",
		"The Role of CDK4/6 Inhibition in Breast Cancer",
		"Mitotic Checkpoint Kinase Mps1 Has a Role in Normal Physiology which Impacts Clinical Utility",
		"Targeting the cyclin-dependent kinases (CDK) 4/6 in estrogen receptor-positive breast cancers",
		"Cyclin-Dependent Kinase Inhibitors as Anticancer Therapeutics",
		"Palbociclib treatment of FLT3-ITD+ AML cells uncovers a kinase-dependent transcriptional regulation of FLT3 and PIM1 by CDK6",
		"Profile of palbociclib in the treatment of metastatic breast cancer",
		"Efficacy of CDK4 inhibition against sarcomas depends on their levels of CDK4 and p16ink4 mRNA",
		"Progress with palbociclib in breast cancer: latest evidence and clinical considerations",
		"Palbociclib: a first-in-class CDK4/CDK6 inhibitor for the treatment of hormone-receptor positive advanced breast cancer",
	}

	c := NewClassifier()

	for _, s := range palbociclib[:len(palbociclib)-2] {
		c.Learn("palbociclib", strings.NewReader(s))
	}
	for _, s := range panitumumab[:len(panitumumab)-2] {
		c.Learn("panitumumab", strings.NewReader(s))
	}

	// for _, s := range c.Categories {
	// 	for w, e := range s.WordFrequencies {
	// 		t.Logf("%20s %d", w, e)
	// 	}
	// }
	t.Log(c.CategoryNames())
	t.Log(c.ProbableCategoreies(strings.NewReader(palbociclib[len(palbociclib)-1])))
	t.Log(c.ProbableCategoreies(strings.NewReader(palbociclib[len(palbociclib)-2])))
	t.Log(c.ProbableCategoreies(strings.NewReader(panitumumab[len(panitumumab)-1])))
	t.Log(c.ProbableCategoreies(strings.NewReader(panitumumab[len(panitumumab)-2])))
}
