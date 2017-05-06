// Kerberos 5 test reference data.
package testdata

const (
	//Expected unmarshaled values
	TEST_REALM                    = "ATHENA.MIT.EDU"
	TEST_CIPHERTEXT               = "krbASN.1 test message"
	TEST_TIME_FORMAT              = "20060102150405"
	TEST_TIME                     = "19940610060317"
	TEST_PRINCIPALNAME_NAMETYPE   = 1
	TEST_KVNO                     = 5
	TEST_ETYPE                    = 0
	TEST_NONCE                    = 42
	TEST_AUTHORIZATION_DATA_TYPE  = 1
	TEST_AUTHORIZATION_DATA_VALUE = "foobar"
	TEST_PADATA_TYPE              = 13
	TEST_PADATA_VALUE             = "pa-data"
)

var TEST_PRINCIPALNAME_NAMESTRING = []string{"hftsai", "extra"}

//The test vectors have been sourced from https://github.com/krb5/krb5/blob/master/src/tests/asn.1/reference_encode.out
var TestVectors = map[string]string{
	"encode_krb5_authenticator":                                  "6281A130819EA003020105A1101B0E415448454E412E4D49542E454455A21A3018A003020101A111300F1B066866747361691B056578747261A30F300DA003020101A106040431323334A405020301E240A511180F31393934303631303036303331375AA6133011A003020101A10A04083132333435363738A703020111A8243022300FA003020101A1080406666F6F626172300FA003020101A1080406666F6F626172",
	"encode_krb5_authenticator(optionalsempty)":                  "624F304DA003020105A1101B0E415448454E412E4D49542E454455A21A3018A003020101A111300F1B066866747361691B056578747261A405020301E240A511180F31393934303631303036303331375A",
	"encode_krb5_authenticator(optionalsNULL)":                   "624F304DA003020105A1101B0E415448454E412E4D49542E454455A21A3018A003020101A111300F1B066866747361691B056578747261A405020301E240A511180F31393934303631303036303331375A",
	"encode_krb5_ticket":                                         "615C305AA003020105A1101B0E415448454E412E4D49542E454455A21A3018A003020101A111300F1B066866747361691B056578747261A3253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765",
	"encode_krb5_keyblock":                                       "3011A003020101A10A04083132333435363738",
	"encode_krb5_enc_tkt_part":                                   "6382011430820110A007030500FEDCBA98A1133011A003020101A10A04083132333435363738A2101B0E415448454E412E4D49542E454455A31A3018A003020101A111300F1B066866747361691B056578747261A42E302CA003020101A12504234544552C4D49542E2C415448454E412E2C57415348494E47544F4E2E4544552C43532EA511180F31393934303631303036303331375AA611180F31393934303631303036303331375AA711180F31393934303631303036303331375AA811180F31393934303631303036303331375AA920301E300DA003020102A106040412D00023300DA003020102A106040412D00023AA243022300FA003020101A1080406666F6F626172300FA003020101A1080406666F6F626172",
	"encode_krb5_enc_tkt_part(optionalsNULL)":                    "6381A53081A2A007030500FEDCBA98A1133011A003020101A10A04083132333435363738A2101B0E415448454E412E4D49542E454455A31A3018A003020101A111300F1B066866747361691B056578747261A42E302CA003020101A12504234544552C4D49542E2C415448454E412E2C57415348494E47544F4E2E4544552C43532EA511180F31393934303631303036303331375AA711180F31393934303631303036303331375A",
	"encode_krb5_enc_kdc_rep_part":                               "7A82010E3082010AA0133011A003020101A10A04083132333435363738A13630343018A0030201FBA111180F31393934303631303036303331375A3018A0030201FBA111180F31393934303631303036303331375AA20302012AA311180F31393934303631303036303331375AA407030500FEDCBA98A511180F31393934303631303036303331375AA611180F31393934303631303036303331375AA711180F31393934303631303036303331375AA811180F31393934303631303036303331375AA9101B0E415448454E412E4D49542E454455AA1A3018A003020101A111300F1B066866747361691B056578747261AB20301E300DA003020102A106040412D00023300DA003020102A106040412D00023",
	"encode_krb5_enc_kdc_rep_part(optionalsNULL)":                "7A81B23081AFA0133011A003020101A10A04083132333435363738A13630343018A0030201FBA111180F31393934303631303036303331375A3018A0030201FBA111180F31393934303631303036303331375AA20302012AA407030500FE5CBA98A511180F31393934303631303036303331375AA711180F31393934303631303036303331375AA9101B0E415448454E412E4D49542E454455AA1A3018A003020101A111300F1B066866747361691B056578747261",
	"encode_krb5_as_rep":                                         "6B81EA3081E7A003020105A10302010BA22630243010A10302010DA209040770612D646174613010A10302010DA209040770612D64617461A3101B0E415448454E412E4D49542E454455A41A3018A003020101A111300F1B066866747361691B056578747261A55E615C305AA003020105A1101B0E415448454E412E4D49542E454455A21A3018A003020101A111300F1B066866747361691B056578747261A3253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765A6253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765",
	"encode_krb5_as_rep(optionalsNULL)":                          "6B81C23081BFA003020105A10302010BA3101B0E415448454E412E4D49542E454455A41A3018A003020101A111300F1B066866747361691B056578747261A55E615C305AA003020105A1101B0E415448454E412E4D49542E454455A21A3018A003020101A111300F1B066866747361691B056578747261A3253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765A6253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765",
	"encode_krb5_tgs_rep":                                        "6D81EA3081E7A003020105A10302010DA22630243010A10302010DA209040770612D646174613010A10302010DA209040770612D64617461A3101B0E415448454E412E4D49542E454455A41A3018A003020101A111300F1B066866747361691B056578747261A55E615C305AA003020105A1101B0E415448454E412E4D49542E454455A21A3018A003020101A111300F1B066866747361691B056578747261A3253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765A6253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765",
	"encode_krb5_tgs_rep(optionalsNULL)":                         "6D81C23081BFA003020105A10302010DA3101B0E415448454E412E4D49542E454455A41A3018A003020101A111300F1B066866747361691B056578747261A55E615C305AA003020105A1101B0E415448454E412E4D49542E454455A21A3018A003020101A111300F1B066866747361691B056578747261A3253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765A6253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765",
	"encode_krb5_ap_req":                                         "6E819D30819AA003020105A10302010EA207030500FEDCBA98A35E615C305AA003020105A1101B0E415448454E412E4D49542E454455A21A3018A003020101A111300F1B066866747361691B056578747261A3253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765A4253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765",
	"encode_krb5_ap_rep":                                         "6F333031A003020105A10302010FA2253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765",
	"encode_krb5_ap_rep_enc_part":                                "7B363034A011180F31393934303631303036303331375AA105020301E240A2133011A003020101A10A04083132333435363738A303020111",
	"encode_krb5_ap_rep_enc_part(optionalsNULL)":                 "7B1C301AA011180F31393934303631303036303331375AA105020301E240",
	"encode_krb5_as_req":                                         "6A8201E4308201E0A103020105A20302010AA32630243010A10302010DA209040770612D646174613010A10302010DA209040770612D64617461A48201AA308201A6A007030500FEDCBA90A11A3018A003020101A111300F1B066866747361691B056578747261A2101B0E415448454E412E4D49542E454455A31A3018A003020101A111300F1B066866747361691B056578747261A411180F31393934303631303036303331375AA511180F31393934303631303036303331375AA611180F31393934303631303036303331375AA70302012AA8083006020100020101A920301E300DA003020102A106040412D00023300DA003020102A106040412D00023AA253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765AB81BF3081BC615C305AA003020105A1101B0E415448454E412E4D49542E454455A21A3018A003020101A111300F1B066866747361691B056578747261A3253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765615C305AA003020105A1101B0E415448454E412E4D49542E454455A21A3018A003020101A111300F1B066866747361691B056578747261A3253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765",
	"encode_krb5_as_req(optionalsNULLexceptsecond_ticket)":       "6A82011430820110A103020105A20302010AA48201023081FFA007030500FEDCBA98A2101B0E415448454E412E4D49542E454455A511180F31393934303631303036303331375AA70302012AA8083006020100020101AB81BF3081BC615C305AA003020105A1101B0E415448454E412E4D49542E454455A21A3018A003020101A111300F1B066866747361691B056578747261A3253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765615C305AA003020105A1101B0E415448454E412E4D49542E454455A21A3018A003020101A111300F1B066866747361691B056578747261A3253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765",
	"encode_krb5_as_req(optionalsNULLexceptserver)":              "6A693067A103020105A20302010AA45B3059A007030500FEDCBA90A2101B0E415448454E412E4D49542E454455A31A3018A003020101A111300F1B066866747361691B056578747261A511180F31393934303631303036303331375AA70302012AA8083006020100020101",
	"encode_krb5_tgs_req":                                        "6C8201E4308201E0A103020105A20302010CA32630243010A10302010DA209040770612D646174613010A10302010DA209040770612D64617461A48201AA308201A6A007030500FEDCBA90A11A3018A003020101A111300F1B066866747361691B056578747261A2101B0E415448454E412E4D49542E454455A31A3018A003020101A111300F1B066866747361691B056578747261A411180F31393934303631303036303331375AA511180F31393934303631303036303331375AA611180F31393934303631303036303331375AA70302012AA8083006020100020101A920301E300DA003020102A106040412D00023300DA003020102A106040412D00023AA253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765AB81BF3081BC615C305AA003020105A1101B0E415448454E412E4D49542E454455A21A3018A003020101A111300F1B066866747361691B056578747261A3253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765615C305AA003020105A1101B0E415448454E412E4D49542E454455A21A3018A003020101A111300F1B066866747361691B056578747261A3253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765",
	"encode_krb5_tgs_req(optionalsNULLexceptsecond_ticket)":      "6C82011430820110A103020105A20302010CA48201023081FFA007030500FEDCBA98A2101B0E415448454E412E4D49542E454455A511180F31393934303631303036303331375AA70302012AA8083006020100020101AB81BF3081BC615C305AA003020105A1101B0E415448454E412E4D49542E454455A21A3018A003020101A111300F1B066866747361691B056578747261A3253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765615C305AA003020105A1101B0E415448454E412E4D49542E454455A21A3018A003020101A111300F1B066866747361691B056578747261A3253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765",
	"encode_krb5_tgs_req(optionalsNULLexceptserver)":             "6C693067A103020105A20302010CA45B3059A007030500FEDCBA90A2101B0E415448454E412E4D49542E454455A31A3018A003020101A111300F1B066866747361691B056578747261A511180F31393934303631303036303331375AA70302012AA8083006020100020101",
	"encode_krb5_kdc_req_body":                                   "308201A6A007030500FEDCBA90A11A3018A003020101A111300F1B066866747361691B056578747261A2101B0E415448454E412E4D49542E454455A31A3018A003020101A111300F1B066866747361691B056578747261A411180F31393934303631303036303331375AA511180F31393934303631303036303331375AA611180F31393934303631303036303331375AA70302012AA8083006020100020101A920301E300DA003020102A106040412D00023300DA003020102A106040412D00023AA253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765AB81BF3081BC615C305AA003020105A1101B0E415448454E412E4D49542E454455A21A3018A003020101A111300F1B066866747361691B056578747261A3253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765615C305AA003020105A1101B0E415448454E412E4D49542E454455A21A3018A003020101A111300F1B066866747361691B056578747261A3253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765",
	"encode_krb5_kdc_req_body(optionalsNULLexceptsecond_ticket)": "3081FFA007030500FEDCBA98A2101B0E415448454E412E4D49542E454455A511180F31393934303631303036303331375AA70302012AA8083006020100020101AB81BF3081BC615C305AA003020105A1101B0E415448454E412E4D49542E454455A21A3018A003020101A111300F1B066866747361691B056578747261A3253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765615C305AA003020105A1101B0E415448454E412E4D49542E454455A21A3018A003020101A111300F1B066866747361691B056578747261A3253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765",
	"encode_krb5_kdc_req_body(optionalsNULLexceptserver)":        "3059A007030500FEDCBA90A2101B0E415448454E412E4D49542E454455A31A3018A003020101A111300F1B066866747361691B056578747261A511180F31393934303631303036303331375AA70302012AA8083006020100020101",
	"encode_krb5_safe":                                           "746E306CA003020105A103020114A24F304DA00A04086B72623564617461A111180F31393934303631303036303331375AA205020301E240A303020111A40F300DA003020102A106040412D00023A50F300DA003020102A106040412D00023A30F300DA003020101A106040431323334",
	"encode_krb5_safe(optionalsNULL)":                            "743E303CA003020105A103020114A21F301DA00A04086B72623564617461A40F300DA003020102A106040412D00023A30F300DA003020101A106040431323334",
	"encode_krb5_priv":                                           "75333031A003020105A103020115A3253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765",
	"encode_krb5_enc_priv_part":                                  "7C4F304DA00A04086B72623564617461A111180F31393934303631303036303331375AA205020301E240A303020111A40F300DA003020102A106040412D00023A50F300DA003020102A106040412D00023",
	"encode_krb5_enc_priv_part(optionalsNULL)":                   "7C1F301DA00A04086B72623564617461A40F300DA003020102A106040412D00023",
	"encode_krb5_cred":                                           "7681F63081F3A003020105A103020116A281BF3081BC615C305AA003020105A1101B0E415448454E412E4D49542E454455A21A3018A003020101A111300F1B066866747361691B056578747261A3253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765615C305AA003020105A1101B0E415448454E412E4D49542E454455A21A3018A003020101A111300F1B066866747361691B056578747261A3253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765A3253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765",
	"encode_krb5_enc_cred_part":                                  "7D8202233082021FA08201DA308201D63081E8A0133011A003020101A10A04083132333435363738A1101B0E415448454E412E4D49542E454455A21A3018A003020101A111300F1B066866747361691B056578747261A307030500FEDCBA98A411180F31393934303631303036303331375AA511180F31393934303631303036303331375AA611180F31393934303631303036303331375AA711180F31393934303631303036303331375AA8101B0E415448454E412E4D49542E454455A91A3018A003020101A111300F1B066866747361691B056578747261AA20301E300DA003020102A106040412D00023300DA003020102A106040412D000233081E8A0133011A003020101A10A04083132333435363738A1101B0E415448454E412E4D49542E454455A21A3018A003020101A111300F1B066866747361691B056578747261A307030500FEDCBA98A411180F31393934303631303036303331375AA511180F31393934303631303036303331375AA611180F31393934303631303036303331375AA711180F31393934303631303036303331375AA8101B0E415448454E412E4D49542E454455A91A3018A003020101A111300F1B066866747361691B056578747261AA20301E300DA003020102A106040412D00023300DA003020102A106040412D00023A10302012AA211180F31393934303631303036303331375AA305020301E240A40F300DA003020102A106040412D00023A50F300DA003020102A106040412D00023",
	"encode_krb5_enc_cred_part(optionalsNULL)":                   "7D82010E3082010AA0820106308201023015A0133011A003020101A10A040831323334353637383081E8A0133011A003020101A10A04083132333435363738A1101B0E415448454E412E4D49542E454455A21A3018A003020101A111300F1B066866747361691B056578747261A307030500FEDCBA98A411180F31393934303631303036303331375AA511180F31393934303631303036303331375AA611180F31393934303631303036303331375AA711180F31393934303631303036303331375AA8101B0E415448454E412E4D49542E454455A91A3018A003020101A111300F1B066866747361691B056578747261AA20301E300DA003020102A106040412D00023300DA003020102A106040412D00023",
	"encode_krb5_error":                                          "7E81BA3081B7A003020105A10302011EA211180F31393934303631303036303331375AA305020301E240A411180F31393934303631303036303331375AA505020301E240A60302013CA7101B0E415448454E412E4D49542E454455A81A3018A003020101A111300F1B066866747361691B056578747261A9101B0E415448454E412E4D49542E454455AA1A3018A003020101A111300F1B066866747361691B056578747261AB0A1B086B72623564617461AC0A04086B72623564617461",
	"encode_krb5_error(optionalsNULL)":                           "7E60305EA003020105A10302011EA305020301E240A411180F31393934303631303036303331375AA505020301E240A60302013CA9101B0E415448454E412E4D49542E454455AA1A3018A003020101A111300F1B066866747361691B056578747261",
	"encode_krb5_authorization_data":                             "3022300FA003020101A1080406666F6F626172300FA003020101A1080406666F6F626172",
	"encode_krb5_padata_sequence":                                "30243010A10302010DA209040770612D646174613010A10302010DA209040770612D64617461",
	"encode_krb5_typed_data":                                     "30243010A00302010DA109040770612D646174613010A00302010DA109040770612D64617461",
	"encode_krb5_padata_sequence(empty)":                         "3000",
	"encode_krb5_etype_info":                                     "30333014A003020100A10D040B4D6F72746F6E27732023303005A0030201013014A003020102A10D040B4D6F72746F6E2773202332",
	"encode_krb5_etype_info(only1)":                              "30163014A003020100A10D040B4D6F72746F6E2773202330",
	"encode_krb5_etype_info(noinfo)":                             "3000",
	"encode_krb5_etype_info2":                                    "3051301EA003020100A10D1B0B4D6F72746F6E2773202330A208040673326B3A2030300FA003020101A208040673326B3A2031301EA003020102A10D1B0B4D6F72746F6E2773202332A208040673326B3A2032",
	"encode_krb5_etype_info2(only1)":                             "3020301EA003020100A10D1B0B4D6F72746F6E2773202330A208040673326B3A2030",
	"encode_krb5_pa_enc_ts":                                      "301AA011180F31393934303631303036303331375AA105020301E240",
	"encode_krb5_pa_enc_ts(nousec)":                              "3013A011180F31393934303631303036303331375A",
	"encode_krb5_enc_data":                                       "3023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765",
	"encode_krb5_enc_data(MSB-setkvno)":                          "3026A003020100A1060204FF000000A21704156B726241534E2E312074657374206D657373616765",
	"encode_krb5_enc_data(kvno= -1)":                             "3023A003020100A1030201FFA21704156B726241534E2E312074657374206D657373616765",
	//"encode_krb5_sam_challenge_2":                                "3022A00D300B04096368616C6C656E6765A111300F300DA003020101A106040431323334",
	//"encode_krb5_sam_challenge_2_body":                           "3064A00302012AA10703050080000000A20B040974797065206E616D65A411040F6368616C6C656E6765206C6162656CA510040E6368616C6C656E67652069707365A6160414726573706F6E73655F70726F6D70742069707365A8050203543210A903020101",
	//"encode_krb5_sam_response_2":                                 "3042A00302012BA10703050080000000A20C040A747261636B2064617461A31D301BA003020101A10402020D36A20E040C6E6F6E6365206F7220736164A4050203543210",
	//"encode_krb5_enc_sam_response_enc_2":                         "301FA003020158A1180416656E635F73616D5F726573706F6E73655F656E635F32",
	//"encode_krb5_pa_for_user":                                    "304BA01A3018A003020101A111300F1B066866747361691B056578747261A1101B0E415448454E412E4D49542E454455A20F300DA003020101A106040431323334A30A1B086B72623564617461",
	//"encode_krb5_pa_s4u_x509_user":                               "3068A0553053A006020400CA149AA11A3018A003020101A111300F1B066866747361691B056578747261A2101B0E415448454E412E4D49542E454455A312041070615F7334755F783530395F75736572A40703050080000000A10F300DA003020101A106040431323334",
	"encode_krb5_ad_kdcissued": "3065A00F300DA003020101A106040431323334A1101B0E415448454E412E4D49542E454455A21A3018A003020101A111300F1B066866747361691B056578747261A3243022300FA003020101A1080406666F6F626172300FA003020101A1080406666F6F626172",
	//"encode_krb5_ad_signedpath_data":                             "3081C7A030302EA01A3018A003020101A111300F1B066866747361691B056578747261A1101B0E415448454E412E4D49542E454455A111180F31393934303631303036303331375AA2323030302EA01A3018A003020101A111300F1B066866747361691B056578747261A1101B0E415448454E412E4D49542E454455A32630243010A10302010DA209040770612D646174613010A10302010DA209040770612D64617461A4243022300FA003020101A1080406666F6F626172300FA003020101A1080406666F6F626172",
	//"encode_krb5_ad_signedpath":                                  "303EA003020101A10F300DA003020101A106040431323334A32630243010A10302010DA209040770612D646174613010A10302010DA209040770612D64617461",
	//"encode_krb5_iakerb_header":                                  "3018A10A04086B72623564617461A20A04086B72623564617461",
	//"encode_krb5_iakerb_finished":                                "3011A10F300DA003020101A106040431323334",
	//"encode_krb5_fast_response":                                  "30819FA02630243010A10302010DA209040770612D646174613010A10302010DA209040770612D64617461A1133011A003020101A10A04083132333435363738A25B3059A011180F31393934303631303036303331375AA105020301E240A2101B0E415448454E412E4D49542E454455A31A3018A003020101A111300F1B066866747361691B056578747261A40F300DA003020101A106040431323334A30302012A",
	//"encode_krb5_pa_fx_fast_reply":                               "A0293027A0253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765",
	//"encode_krb5_otp_tokeninfo(optionalsNULL)":                   "300780050000000000",
	//"encode_krb5_otp_tokeninfo":                                  "307280050077000000810B4578616D706C65636F727082056861726B2183010A8401028509796F7572746F6B656E862875726E3A696574663A706172616D733A786D6C3A6E733A6B657970726F763A70736B633A686F7470A716300B0609608648016503040201300706052B0E03021A880203E8",
	//"encode_krb5_pa_otp_challenge(optionalsNULL)":                "301580086D696E6E6F6E6365A209300780050000000000",
	//"encode_krb5_pa_otp_challenge":                               "3081A580086D61786E6F6E6365810B7465737473657276696365A27D300780050000000000307280050077000000810B4578616D706C65636F727082056861726B2183010A8401028509796F7572746F6B656E862875726E3A696574663A706172616D733A786D6C3A6E733A6B657970726F763A70736B633A686F7470A716300B0609608648016503040201300706052B0E03021A880203E883076B657973616C74840431323334",
	//"encode_krb5_pa_otp_req(optionalsNULL)":                      "302C80050000000000A223A003020100A103020105A21704156B726241534E2E312074657374206D657373616765",
	//"encode_krb5_pa_otp_req":                                     "3081B98005006000000081056E6F6E6365A223A003020100A103020105A21704156B726241534E2E312074657374206D657373616765A30B0609608648016503040201840203E8850566726F6773860A6D79666972737470696E87056861726B21880F31393934303631303036303331375A89033334368A01028B09796F7572746F6B656E8C2875726E3A696574663A706172616D733A786D6C3A6E733A6B657970726F763A70736B633A686F74708D0B4578616D706C65636F7270",
	//"encode_krb5_pa_otp_enc_req":                                 "300A80086B72623564617461",
	//"encode_krb5_kkdcp_message":                                  "308201FCA08201EC048201E86A8201E4308201E0A103020105A20302010AA32630243010A10302010DA209040770612D646174613010A10302010DA209040770612D64617461A48201AA308201A6A007030500FEDCBA98A11A3018A003020101A111300F1B066866747361691B056578747261A2101B0E415448454E412E4D49542E454455A31A3018A003020101A111300F1B066866747361691B056578747261A411180F31393934303631303036303331375AA511180F31393934303631303036303331375AA611180F31393934303631303036303331375AA70302012AA8083006020100020101A920301E300DA003020102A106040412D00023300DA003020102A106040412D00023AA253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765AB81BF3081BC615C305AA003020105A1101B0E415448454E412E4D49542E454455A21A3018A003020101A111300F1B066866747361691B056578747261A3253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765615C305AA003020105A1101B0E415448454E412E4D49542E454455A21A3018A003020101A111300F1B066866747361691B056578747261A3253023A003020100A103020105A21704156B726241534E2E312074657374206D657373616765A10A1B086B72623564617461",
	//"encode_krb5_cammac(optionalsNULL)":                          "3012A010300E300CA003020101A1050403616431",
	//"encode_krb5_cammac":                                         "3081F2A01E301C300CA003020101A1050403616431300CA003020102A1050403616432A13D303BA01A3018A003020101A111300F1B066866747361691B056578747261A103020105A203020110A3133011A003020101A10A0408636B73756D6B6463A23D303BA01A3018A003020101A111300F1B066866747361691B056578747261A103020105A203020110A3133011A003020101A10A0408636B73756D737663A35230503013A311300FA003020101A1080406636B73756D313039A01A3018A003020101A111300F1B066866747361691B056578747261A103020105A203020110A311300FA003020101A1080406636B73756D32",
	//"encode_krb5_secure_cookie":                                  "302C02042DF8022530243010A10302010DA209040770612D646174613010A10302010DA209040770612D64617461",

	//MS PAC Vectors
	"PAC_AuthorizationData": "308205523082054ea00402020080a182054404820540040000000000000001000000b004000048000000000000000a00000012000000f804000000000000060000001400000010050000000000000700000014000000280500000000000001100800cccccccca00400000000000000000200d186660f656ac601ffffffffffffff7fffffffffffffff7f17d439fe784ac6011794a328424bc601175424977a81c60108000800040002002400240008000200120012000c0002000000000010000200000000001400020000000000180002005410000097792c00010200001a0000001c000200200000000000000000000000000000000000000016001800200002000a000c002400020028000200000000000000000010000000000000000000000000000000000000000000000000000000000000000d0000002c0002000000000000000000000000000400000000000000040000006c007a00680075001200000000000000120000004c0069007100690061006e00670028004c006100720072007900290020005a00680075000900000000000000090000006e0074006400730032002e0062006100740000000000000000000000000000000000000000000000000000000000000000000000000000001a00000061c433000700000009c32d00070000005eb4320007000000010200000700000097b92c00070000002bf1320007000000ce30330007000000a72e2e00070000002af132000700000098b92c000700000062c4330007000000940133000700000076c4330007000000aefe2d000700000032d22c00070000001608320007000000425b2e00070000005fb4320007000000ca9c35000700000085442d0007000000c2f0320007000000e9ea310007000000ed8e2e0007000000b6eb310007000000ab2e2e0007000000720e2e00070000000c000000000000000b0000004e0054004400450056002d00440043002d003000350000000600000000000000050000004e0054004400450056000000040000000104000000000005150000005951b81766725d2564633b0b0d0000003000020007000000340002000700002038000200070000203c000200070000204000020007000020440002000700002048000200070000204c000200070000205000020007000020540002000700002058000200070000205c00020007000020600002000700002005000000010500000000000515000000b9301b2eb7414c6c8c3b351501020000050000000105000000000005150000005951b81766725d2564633b0b74542f00050000000105000000000005150000005951b81766725d2564633b0be8383200050000000105000000000005150000005951b81766725d2564633b0bcd383200050000000105000000000005150000005951b81766725d2564633b0b5db43200050000000105000000000005150000005951b81766725d2564633b0b41163500050000000105000000000005150000005951b81766725d2564633b0be8ea3100050000000105000000000005150000005951b81766725d2564633b0bc1193200050000000105000000000005150000005951b81766725d2564633b0b29f13200050000000105000000000005150000005951b81766725d2564633b0b0f5f2e00050000000105000000000005150000005951b81766725d2564633b0b2f5b2e00050000000105000000000005150000005951b81766725d2564633b0bef8f3100050000000105000000000005150000005951b81766725d2564633b0b075f2e00000000000049d90e656ac60108006c007a006800750000000000000076ffffff41edce9a34815d3aef7bc98874805d250000000076fffffff7a534dab2c02986efe0fbe5110a4f3200000000",
}

const (
	TESTUSER1_KEYTAB      = "05020000003b0001000b544553542e474f4b52423500097465737475736572310000000158ee757c0100110010698c4df8e9f60e7eea5a21bf4526ad25000000010000004b0001000b544553542e474f4b52423500097465737475736572310000000158ee757c0100120020bbdc430aab7e2d4622a0b6951481453b0962e9db8e2f168942ad175cda6d9de900000001"
	TESTUSER1_WRONGPASSWD = "0502000000370001000b544553542e474f4b52423500097465737475736572310000000158ef4bc5010011001039a9a382153105f8708e80f93382654e000000470001000b544553542e474f4b52423500097465737475736572310000000158ef4bc60100120020fc5bb940d6075214e0c6fc0456ce68c33306094198a927b4187d7cf3f4aea50d"
	TESTUSER2_KEYTAB      = "05020000003b0001000b544553542e474f4b52423500097465737475736572320000000158ee75db010011001086824c55ff5de30386dd83dc62b44bb7000000010000004b0001000b544553542e474f4b52423500097465737475736572320000000158ee75db0100120020d8ed27f96be76fd5b281ee9f8029db93cc5fb06c7eb3be9ee753106d3488fa9200000001"
	TESTUSER3_KEYTAB      = "05020000003b0001000b544553542e474f4b52423500097465737475736572330000000158ee76010100110010220789f5cf68e44a852c73b1e3729efc000000010000004b0001000b544553542e474f4b52423500097465737475736572330000000158ee760101001200205e1b7287f78f3f88c9eefe92c08ec1613f35d72bc6ee2c2e1c5cb6fb9d8b058000000001"
	HTTP_KEYTAB           = "0502000000480002000b544553542e474f4b5242350004485454500010686f73742e746573742e676f6b7262350000000158ee7636010011001057a7754c70c4d85c155c718c2f1292b000000001000000580002000b544553542e474f4b5242350004485454500010686f73742e746573742e676f6b7262350000000158ee763601001200209cad00bbc72d703258e911dc18e6d5487cf737bf67fd111f0c2463ad6033bf5100000001"
	TEST_AS_REQ           = "6a81a63081a3a103020105a20302010aa30e300c300aa10402020095a2020400a48186308183a00703050040000010a1163014a003020101a10d300b1b09746573747573657231a20d1b0b544553542e474f4b524235a320301ea003020102a11730151b066b72627467741b0b544553542e474f4b524235a511180f32303137303232303134323530315aa70602040f6755a6a814301202011202011102011002011702011902011a"
	TEST_AS_REP           = "6b8202f3308202efa003020105a10302010ba22e302c302aa103020113a2230421301f301da003020112a1161b14544553542e474f4b524235746573747573657231a30d1b0b544553542e474f4b524235a4163014a003020101a10d300b1b09746573747573657231a582015a6182015630820152a003020105a10d1b0b544553542e474f4b524235a220301ea003020102a11730151b066b72627467741b0b544553542e474f4b524235a382011830820114a003020112a103020101a28201060482010264d3fa49d89b627ed471298846ff92cd8632f657c58fe25322a61fffa32bb7966dc4c44c86a81353def2a11c36c537191406a609147f424a63266c00d02bcc56a27b0969d86ff4352634be9e2a4ac0ad5a36b0b0a3d689f128c0afa97401796e88037a35ad19efaf31d1ed4f3213769c03a58bc90ffac2051db152c0ed0809ad05ffb03aa3afaf731ed85f7a73020cb72355e0de27842dcf7eae3de9f7c14aa237edb25153b217ef3693373bc3cacbebe406910ff9ae9d00b7b08f726cb29a213cb9ad51ba80a8c24fa4b6692a445686889702cfa6ea749bac03e27e982407aca623fbd48586bcf566cfe87e1d9f17a74b1315669c16480f93e9d8782e71a8f11000a682012c30820128a003020112a282011f0482011b99b86153c0393c0e4130628f3e1e0f0a1f034e7e61a111b7fad15884e231c8fd8727e0bc945c9b35be20c57d057c8b09b0de74c53fb38cc15c9a2d483023fc369f5bde4da7324b4732b5a3d9504d92f67026aaa01df4f0138245d2ccb1c5a4014804cf295c7e7e56a867e6cf0c534f667f32da7aa5e700af1461764f1c276a8ff0fbee0e99322fe2059d2321853be09d0956c3afcfd07e3e702646a4678926a77bea20d9aaf3086b6d384821c81900af9013a3519f0e50eab6e1491d72e4ee17c2a44441b2ebc8a796cc3d876e328347dce65f61104e14d4c31532885776c9c8a70186b8b39f928972945c98bd60381ead5448e7ebe93fea308054287ac34b0583b4b9b5e43c5f8518d693ba9eb48a219c27344466b3c693a70462"
	TEST_TGS_REQ          = "6c82038f3082038ba103020105a20302010ca382031a3082031630820245a103020101a282023c048202386e82023430820230a003020105a10302010ea20703050000000000a382015a6182015630820152a003020105a10d1b0b544553542e474f4b524235a220301ea003020102a11730151b066b72627467741b0b544553542e474f4b524235a382011830820114a003020112a103020101a28201060482010264d3fa49d89b627ed471298846ff92cd8632f657c58fe25322a61fffa32bb7966dc4c44c86a81353def2a11c36c537191406a609147f424a63266c00d02bcc56a27b0969d86ff4352634be9e2a4ac0ad5a36b0b0a3d689f128c0afa97401796e88037a35ad19efaf31d1ed4f3213769c03a58bc90ffac2051db152c0ed0809ad05ffb03aa3afaf731ed85f7a73020cb72355e0de27842dcf7eae3de9f7c14aa237edb25153b217ef3693373bc3cacbebe406910ff9ae9d00b7b08f726cb29a213cb9ad51ba80a8c24fa4b6692a445686889702cfa6ea749bac03e27e982407aca623fbd48586bcf566cfe87e1d9f17a74b1315669c16480f93e9d8782e71a8f11000a481bc3081b9a003020112a281b10481ae8ae3cb8ac47d77cfc7b0b6bf0d3c5f8fcc6dd569344256a6a40c004fc2d23ebbe6ee0b9e00eccf37e710b7c01a7d2a63bbed6d75f2b230d24d724ef90edad2c5680e7e2436ab1145ff68481673444ebd61e3aef79b9ee05809551672c6c436eb8ac732a7fe78bd8f380e68a541191e3125554e4bab63dcc19ea931c1477366a6039ff7b7e62521ebfeffd6784b6ef0c97f653ac4d8dfb304f3e2e843faab12d838c23f1105f0a281c39325987cb03081caa10402020088a281c10481bea081bb3081b8a1173015a003020110a10e040ce613d8e9d544f0e56c60d3bba2819c308199a003020112a2819104818ec4fabcb1ec2f24e04ef51f9247239b28275653fa5cbc1dc9e747530c597631050fe86a5f3cba2ff54270aa771dcefa87efc8c8604407f84e603f5c01a2d929e18103561c3ffbc3a0cf63340bdd67a0739d4d81989827fc1d3f7f13e9dd5cc2346ca08e26a2aaf6d0102fbef8f7a6ee0a1caae7880e953ea678da619038786122a0b71853e8d0b95f544f8fbd6945a461305fa00703050040810000a20d1b0b544553542e474f4b524235a3233021a003020101a11a30181b04485454501b10686f73742e746573742e676f6b726235a511180f32303137303232303032323634325aa706020458a9ab2aa8053003020112"
	TEST_TGS_REP          = "6d82039d30820399a003020105a10302010da281df3081dc3081d9a10402020088a281d00481cda081ca3081c7a081c43081c1a003020112a281b90481b62e2b31bd998e2747a447175bcdbe14b7bc043c747c3289e10cde7403ee685d64f9724055c667ad2d3ce6240b5457ddfe3505a5b4d5d5e35238522ba5d86f1b691f28bbe7290487ffebf7e720ecf772cef061ac2c433cabbd0d5973fbedeb5418c8ba1bb7149dd625a4ff4cb465d9599c9d4cfb899807fa088521fa0e0ed5cfa4dd6d02b2a6854c4feb4ec382de5820134ca7c140be4b0f416ca5ebb328c8a470a55154062b8070683d0897a40277bdc752e94437920aa30d1b0b544553542e474f4b524235a4163014a003020101a10d300b1b09746573747573657231a58201706182016c30820168a003020105a10d1b0b544553542e474f4b524235a2233021a003020101a11a30181b04485454501b10686f73742e746573742e676f6b726235a382012b30820127a003020112a103020102a28201190482011524db012d81dde2cb3b7b40a35ce4fe17f7898166c7a7534ad73ca761e75316a06504415b1bc62508edb6ea58544a3ac0f98911cd55332442d30e007e0f39efa939b726a5228514984a133d1ba7d93d60fec5b2f7fbb16356c85ad1b0bb7ed6108420514086ea9959037e794b535fa052651385060da83f93acffbde0cb486b2f236f8955bb521ef90bc0a52944d2a8da82389e6861065064cc5178a5a7d302e9950761648726fc4015f8772339ef401fd8327dd335c6692c010f2c31b5ccccbcd83e68f2a1b66f16ca44e0bfb2903801c27ed8dd9a7b2dcea3c2bc91d8055c603f218494b1342490fbb805492d999491c2570b3ab392ba7e62c23659663509838e91b6f560a284889075071c348c701017a0a5e5c7a682010e3082010aa003020112a28201010481fece77ad9ebec9b094b686d4885895a2530471e9a56a93c2a463eaa988932695d77a12b6c02e69edfea6759ed49b029bdef43e446129b80a1ef95b9cea69a382d5c074868e68b676869f0990cad588c6d3589b9bba795ff0572ba648915f68a3b49df1def285645e93f4285a843ead01122c4f00bebefa33ecf6df2584e33da359d3d53e0a9b8ec0f41120a0d612d3a1e4504ce464dddec5e2aa3f0d2b6dae7b8800091760d7a590eaa6d7b33539c1700d9daf22a7c8cc1e108aee75267434997666132f52c08ea7c0c898859214ac4e71e185d3bcec1467ab6e8d91e79ba2b2041d40714da495c79d232d3c4e72a0910b9143fd648fe24bc46ec3416a3c34"
	TEST_KRB5CONF         = `[libdefaults]
  default_realm = TEST.GOKRB5
  dns_lookup_realm = false
  dns_lookup_kdc = false
  ticket_lifetime = 24h
  forwardable = yes
  default_tkt_enctypes = aes256-cts-hmac-sha1-96
  default_tgs_enctypes = aes256-cts-hmac-sha1-96

[realms]
 TEST.GOKRB5 = {
  kdc = 10.80.88.88:88
  admin_server = 10.80.88.88:749
  default_domain = test.gokrb5
 }

[domain_realm]
 .test.gokrb5 = TEST.GOKRB5
 test.gokrb5 = TEST.GOKRB5
 `
	TEST_KRB5CONF_TCP = `[libdefaults]
  default_realm = TEST.GOKRB5
  dns_lookup_realm = false
  dns_lookup_kdc = false
  ticket_lifetime = 24h
  forwardable = yes
  default_tkt_enctypes = aes256-cts-hmac-sha1-96
  default_tgs_enctypes = aes256-cts-hmac-sha1-96
  udp_preference_limit = 1

[realms]
 TEST.GOKRB5 = {
  kdc = 10.80.88.88:88
  admin_server = 10.80.88.88:749
  default_domain = test.gokrb5
 }

[domain_realm]
 .test.gokrb5 = TEST.GOKRB5
 test.gokrb5 = TEST.GOKRB5
 `
	TEST_KRB5CONF_BAD_KDC_ADDRESS = `[libdefaults]
  default_realm = TEST.GOKRB5
  dns_lookup_realm = false
  dns_lookup_kdc = false
  ticket_lifetime = 24h
  forwardable = yes
  default_tkt_enctypes = aes256-cts-hmac-sha1-96
  default_tgs_enctypes = aes256-cts-hmac-sha1-96

[realms]
 TEST.GOKRB5 = {
  kdc = 10.80.88.153:88
  admin_server = 10.80.88.153:749
  default_domain = test.gokrb5
 }

[domain_realm]
 .test.gokrb5 = TEST.GOKRB5
 test.gokrb5 = TEST.GOKRB5
 `
	TEST_KRB5CONF_OLDERKDC = `[libdefaults]
  default_realm = TEST.GOKRB5
  dns_lookup_realm = false
  dns_lookup_kdc = false
  ticket_lifetime = 24h
  forwardable = yes
  default_tkt_enctypes = aes256-cts-hmac-sha1-96
  default_tgs_enctypes = aes256-cts-hmac-sha1-96

[realms]
 TEST.GOKRB5 = {
  kdc = 10.80.88.78:88
  admin_server = 10.80.88.78:749
  default_domain = test.gokrb5
 }

[domain_realm]
 .test.gokrb5 = TEST.GOKRB5
 test.gokrb5 = TEST.GOKRB5
 `
	TEST_KRB5CONF_AD = `[libdefaults]
  default_realm = TEST.GOKRB5
  dns_lookup_realm = false
  dns_lookup_kdc = false
  ticket_lifetime = 24h
  forwardable = yes
  default_tkt_enctypes = aes256-cts-hmac-sha1-96
  default_tgs_enctypes = aes256-cts-hmac-sha1-96
  udp_preference_limit = 1

[realms]
 TEST.GOKRB5 = {
  kdc = 10.80.88.68:88
  admin_server = 10.80.88.68:749
  default_domain = test.gokrb5
 }

[domain_realm]
 .test.gokrb5 = TEST.GOKRB5
 test.gokrb5 = TEST.GOKRB5
 `
)
