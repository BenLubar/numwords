package numwords_test

import (
	"fmt"
	"math/big"

	"github.com/BenLubar/numwords"
)

func ExampleInt() {
	fmt.Println(numwords.Int(123456789))

	// Output:
	// one hundred twenty three million four hundred fifty six thousand seven hundred eighty nine
}

func ExampleBig() {
	var n big.Int
	n.SetString("795000251651046142383363615387133118749508854110191293727347242043532350777196212610827095586063001495724835917929521451989134919276118489572386025458987675560466237687214202228679855843675243516087540059396364947554108055190618240585503850322746581731927830120842751492094276657306374411113989025103065871065954416297675507346487666715477206256039879117492991185564560158455309667201974686829457745359047768182328742319895928966044805750709148051363297693272806208168391852100263601306836475323016696941146887126686704671110002824057711375130894433491574005685227138528967429782474487029487409502135877073585102515450654789458996778668276134128641623044490707747290142790354677000402083905237020713023176559895738857691020219128514393867041592808036334097045828988769202834691155136934149649927954069517836448083618260387293658864280162736381004779808990547017089811459827643947342670481652870383619946898992579756003415775059034397544086130996241320266097099381600900861663763260272863517947006272470516300671897902287446644461522463944391638270486597168622580456460184346362437392520975362854756293708534931555419438096726622538931993299672573631512774359574591677712288145813008338560114244784810192622670367526590124947188470737", 10)

	fmt.Println(numwords.Big(&n))

	// Output:
	// seven hundred ninety five novenquadringentiillion two hundred fifty one septenquadringentiillion six hundred fifty one sequadringentiillion forty six quinquaquadringentiillion one hundred forty two quattuorquadringentiillion three hundred eighty three tresquadringentiillion three hundred sixty three duoquadringentiillion six hundred fifteen unquadringentiillion three hundred eighty seven quadringentiillion one hundred thirty three novenonagintatrecentiillion one hundred eighteen octononagintatrecentiillion seven hundred forty nine septenonagintatrecentiillion five hundred eight senonagintatrecentiillion eight hundred fifty four quinquanonagintatrecentiillion one hundred ten quattuornonagintatrecentiillion one hundred ninety one trenonagintatrecentiillion two hundred ninety three duononagintatrecentiillion seven hundred twenty seven unnonagintatrecentiillion three hundred forty seven nonagintatrecentiillion two hundred forty two novemoctogintatrecentiillion forty three octooctogintatrecentiillion five hundred thirty two septemoctogintatrecentiillion three hundred fifty sexoctogintatrecentiillion seven hundred seventy seven quinquaoctogintatrecentiillion one hundred ninety six quattuoroctogintatrecentiillion two hundred twelve treoctogintatrecentiillion six hundred ten duooctogintatrecentiillion eight hundred twenty seven unoctogintatrecentiillion ninety five octogintatrecentiillion five hundred eighty six novenseptuagintatrecentiillion sixty three octoseptuagintatrecentiillion one septenseptuagintatrecentiillion four hundred ninety five seseptuagintatrecentiillion seven hundred twenty four quinquaseptuagintatrecentiillion eight hundred thirty five quattuorseptuagintatrecentiillion nine hundred seventeen treseptuagintatrecentiillion nine hundred twenty nine duoseptuagintatrecentiillion five hundred twenty one unseptuagintatrecentiillion four hundred fifty one septuagintatrecentiillion nine hundred eighty nine novensexagintatrecentiillion one hundred thirty four octosexagintatrecentiillion nine hundred nineteen septensexagintatrecentiillion two hundred seventy six sesexagintatrecentiillion one hundred eighteen quinquasexagintatrecentiillion four hundred eighty nine quattuorsexagintatrecentiillion five hundred seventy two tresexagintatrecentiillion three hundred eighty six duosexagintatrecentiillion twenty five unsexagintatrecentiillion four hundred fifty eight sexagintatrecentiillion nine hundred eighty seven novenquinquagintatrecentiillion six hundred seventy five octoquinquagintatrecentiillion five hundred sixty septenquinquagintatrecentiillion four hundred sixty six sequinquagintatrecentiillion two hundred thirty seven quinquaquinquagintatrecentiillion six hundred eighty seven quattuorquinquagintatrecentiillion two hundred fourteen tresquinquagintatrecentiillion two hundred two duoquinquagintatrecentiillion two hundred twenty eight unquinquagintatrecentiillion six hundred seventy nine quinquagintatrecentiillion eight hundred fifty five novenquadragintatrecentiillion eight hundred forty three octoquadragintatrecentiillion six hundred seventy five septenquadragintatrecentiillion two hundred forty three sequadragintatrecentiillion five hundred sixteen quinquaquadragintatrecentiillion eighty seven quattuorquadragintatrecentiillion five hundred forty tresquadragintatrecentiillion fifty nine duoquadragintatrecentiillion three hundred ninety six unquadragintatrecentiillion three hundred sixty four quadragintatrecentiillion nine hundred forty seven noventrigintatrecentiillion five hundred fifty four octotrigintatrecentiillion one hundred eight septentrigintatrecentiillion fifty five setrigintatrecentiillion one hundred ninety quinquatrigintatrecentiillion six hundred eighteen quattuortrigintatrecentiillion two hundred forty trestrigintatrecentiillion five hundred eighty five duotrigintatrecentiillion five hundred three untrigintatrecentiillion eight hundred fifty trigintatrecentiillion three hundred twenty two novemvigintitrecentiillion seven hundred forty six octovigintitrecentiillion five hundred eighty one septemvigintitrecentiillion seven hundred thirty one sevigintitrecentiillion nine hundred twenty seven quinquavigintitrecentiillion eight hundred thirty quattuorvigintitrecentiillion one hundred twenty tresvigintitrecentiillion eight hundred forty two duovigintitrecentiillion seven hundred fifty one unvigintitrecentiillion four hundred ninety two vigintitrecentiillion ninety four novendecitrecentiillion two hundred seventy six octodecitrecentiillion six hundred fifty seven septendecitrecentiillion three hundred six sedecitrecentiillion three hundred seventy four quinquadecitrecentiillion four hundred eleven quattuordecitrecentiillion one hundred thirteen tredecitrecentiillion nine hundred eighty nine duodecitrecentiillion twenty five undecitrecentiillion one hundred three decitrecentiillion sixty five noventrecentiillion eight hundred seventy one octotrecentiillion sixty five septentrecentiillion nine hundred fifty four setrecentiillion four hundred sixteen quinquatrecentiillion two hundred ninety seven quattuortrecentiillion six hundred seventy five trestrecentiillion five hundred seven duotrecentiillion three hundred forty six untrecentiillion four hundred eighty seven trecentiillion six hundred sixty six novenonagintaducentiillion seven hundred fifteen octononagintaducentiillion four hundred seventy seven septenonagintaducentiillion two hundred six senonagintaducentiillion two hundred fifty six quinquanonagintaducentiillion thirty nine quattuornonagintaducentiillion eight hundred seventy nine trenonagintaducentiillion one hundred seventeen duononagintaducentiillion four hundred ninety two unnonagintaducentiillion nine hundred ninety one nonagintaducentiillion one hundred eighty five novemoctogintaducentiillion five hundred sixty four octooctogintaducentiillion five hundred sixty septemoctogintaducentiillion one hundred fifty eight sexoctogintaducentiillion four hundred fifty five quinquaoctogintaducentiillion three hundred nine quattuoroctogintaducentiillion six hundred sixty seven treoctogintaducentiillion two hundred one duooctogintaducentiillion nine hundred seventy four unoctogintaducentiillion six hundred eighty six octogintaducentiillion eight hundred twenty nine novenseptuagintaducentiillion four hundred fifty seven octoseptuagintaducentiillion seven hundred forty five septenseptuagintaducentiillion three hundred fifty nine seseptuagintaducentiillion forty seven quinquaseptuagintaducentiillion seven hundred sixty eight quattuorseptuagintaducentiillion one hundred eighty two treseptuagintaducentiillion three hundred twenty eight duoseptuagintaducentiillion seven hundred forty two unseptuagintaducentiillion three hundred nineteen septuagintaducentiillion eight hundred ninety five novensexagintaducentiillion nine hundred twenty eight octosexagintaducentiillion nine hundred sixty six septensexagintaducentiillion forty four sesexagintaducentiillion eight hundred five quinquasexagintaducentiillion seven hundred fifty quattuorsexagintaducentiillion seven hundred nine tresexagintaducentiillion one hundred forty eight duosexagintaducentiillion fifty one unsexagintaducentiillion three hundred sixty three sexagintaducentiillion two hundred ninety seven novenquinquagintaducentiillion six hundred ninety three octoquinquagintaducentiillion two hundred seventy two septenquinquagintaducentiillion eight hundred six sequinquagintaducentiillion two hundred eight quinquaquinquagintaducentiillion one hundred sixty eight quattuorquinquagintaducentiillion three hundred ninety one tresquinquagintaducentiillion eight hundred fifty two duoquinquagintaducentiillion one hundred unquinquagintaducentiillion two hundred sixty three quinquagintaducentiillion six hundred one novenquadragintaducentiillion three hundred six octoquadragintaducentiillion eight hundred thirty six septenquadragintaducentiillion four hundred seventy five sequadragintaducentiillion three hundred twenty three quinquaquadragintaducentiillion sixteen quattuorquadragintaducentiillion six hundred ninety six tresquadragintaducentiillion nine hundred forty one duoquadragintaducentiillion one hundred forty six unquadragintaducentiillion eight hundred eighty seven quadragintaducentiillion one hundred twenty six noventrigintaducentiillion six hundred eighty six octotrigintaducentiillion seven hundred four septentrigintaducentiillion six hundred seventy one setrigintaducentiillion one hundred ten quinquatrigintaducentiillion two quattuortrigintaducentiillion eight hundred twenty four trestrigintaducentiillion fifty seven duotrigintaducentiillion seven hundred eleven untrigintaducentiillion three hundred seventy five trigintaducentiillion one hundred thirty novemvigintiducentiillion eight hundred ninety four octovigintiducentiillion four hundred thirty three septemvigintiducentiillion four hundred ninety one sevigintiducentiillion five hundred seventy four quinquavigintiducentiillion five quattuorvigintiducentiillion six hundred eighty five tresvigintiducentiillion two hundred twenty seven duovigintiducentiillion one hundred thirty eight unvigintiducentiillion five hundred twenty eight vigintiducentiillion nine hundred sixty seven novendeciducentiillion four hundred twenty nine octodeciducentiillion seven hundred eighty two septendeciducentiillion four hundred seventy four sedeciducentiillion four hundred eighty seven quinquadeciducentiillion twenty nine quattuordeciducentiillion four hundred eighty seven tredeciducentiillion four hundred nine duodeciducentiillion five hundred two undeciducentiillion one hundred thirty five deciducentiillion eight hundred seventy seven novenducentiillion seventy three octoducentiillion five hundred eighty five septenducentiillion one hundred two seducentiillion five hundred fifteen quinquaducentiillion four hundred fifty quattuorducentiillion six hundred fifty four treducentiillion seven hundred eighty nine duoducentiillion four hundred fifty eight unducentiillion nine hundred ninety six ducentiillion seven hundred seventy eight novenonagintacentiillion six hundred sixty eight octononagintacentiillion two hundred seventy six septenonagintacentiillion one hundred thirty four senonagintacentiillion one hundred twenty eight quinquanonagintacentiillion six hundred forty one quattuornonagintacentiillion six hundred twenty three trenonagintacentiillion forty four duononagintacentiillion four hundred ninety unnonagintacentiillion seven hundred seven nonagintacentiillion seven hundred forty seven novemoctogintacentiillion two hundred ninety octooctogintacentiillion one hundred forty two septemoctogintacentiillion seven hundred ninety sexoctogintacentiillion three hundred fifty four quinquaoctogintacentiillion six hundred seventy seven quattuoroctogintacentiillion four hundred two duooctogintacentiillion eighty three unoctogintacentiillion nine hundred five octogintacentiillion two hundred thirty seven novenseptuagintacentiillion twenty octoseptuagintacentiillion seven hundred thirteen septenseptuagintacentiillion twenty three seseptuagintacentiillion one hundred seventy six quinquaseptuagintacentiillion five hundred fifty nine quattuorseptuagintacentiillion eight hundred ninety five treseptuagintacentiillion seven hundred thirty eight duoseptuagintacentiillion eight hundred fifty seven unseptuagintacentiillion six hundred ninety one septuagintacentiillion twenty novensexagintacentiillion two hundred nineteen octosexagintacentiillion one hundred twenty eight septensexagintacentiillion five hundred fourteen sesexagintacentiillion three hundred ninety three quinquasexagintacentiillion eight hundred sixty seven quattuorsexagintacentiillion forty one tresexagintacentiillion five hundred ninety two duosexagintacentiillion eight hundred eight unsexagintacentiillion thirty six sexagintacentiillion three hundred thirty four novenquinquagintacentiillion ninety seven octoquinquagintacentiillion forty five septenquinquagintacentiillion eight hundred twenty eight sequinquagintacentiillion nine hundred eighty eight quinquaquinquagintacentiillion seven hundred sixty nine quattuorquinquagintacentiillion two hundred two tresquinquagintacentiillion eight hundred thirty four duoquinquagintacentiillion six hundred ninety one unquinquagintacentiillion one hundred fifty five quinquagintacentiillion one hundred thirty six novenquadragintacentiillion nine hundred thirty four octoquadragintacentiillion one hundred forty nine septenquadragintacentiillion six hundred forty nine sequadragintacentiillion nine hundred twenty seven quinquaquadragintacentiillion nine hundred fifty four quattuorquadragintacentiillion sixty nine tresquadragintacentiillion five hundred seventeen duoquadragintacentiillion eight hundred thirty six unquadragintacentiillion four hundred forty eight quadragintacentiillion eighty three noventrigintacentiillion six hundred eighteen octotrigintacentiillion two hundred sixty septentrigintacentiillion three hundred eighty seven setrigintacentiillion two hundred ninety three quinquatrigintacentiillion six hundred fifty eight quattuortrigintacentiillion eight hundred sixty four trestrigintacentiillion two hundred eighty duotrigintacentiillion one hundred sixty two untrigintacentiillion seven hundred thirty six trigintacentiillion three hundred eighty one novemviginticentiillion four octoviginticentiillion seven hundred seventy nine septemviginticentiillion eight hundred eight seviginticentiillion nine hundred ninety quinquaviginticentiillion five hundred forty seven quattuorviginticentiillion seventeen tresviginticentiillion eighty nine duoviginticentiillion eight hundred eleven unviginticentiillion four hundred fifty nine viginticentiillion eight hundred twenty seven novendecicentiillion six hundred forty three octodecicentiillion nine hundred forty seven septendecicentiillion three hundred forty two sedecicentiillion six hundred seventy quinquadecicentiillion four hundred eighty one quattuordecicentiillion six hundred fifty two tredecicentiillion eight hundred seventy duodecicentiillion three hundred eighty three undecicentiillion six hundred nineteen decicentiillion nine hundred forty six novencentiillion eight hundred ninety eight octocentiillion nine hundred ninety two septencentiillion five hundred seventy nine sexcentiillion seven hundred fifty six quinquacentiillion three quattuorcentiillion four hundred fifteen trecentiillion seven hundred seventy five duocentiillion fifty nine uncentiillion thirty four centiillion three hundred ninety seven novenonagintaillion five hundred forty four octononagintaillion eighty six septenonagintaillion one hundred thirty senonagintaillion nine hundred ninety six quinquanonagintaillion two hundred forty one quattuornonagintaillion three hundred twenty trenonagintaillion two hundred sixty six duononagintaillion ninety seven unnonagintaillion ninety nine nonagintaillion three hundred eighty one novemoctogintaillion six hundred octooctogintaillion nine hundred septemoctogintaillion eight hundred sixty one sexoctogintaillion six hundred sixty three quinquaoctogintaillion seven hundred sixty three quattuoroctogintaillion two hundred sixty treoctogintaillion two hundred seventy two duooctogintaillion eight hundred sixty three unoctogintaillion five hundred seventeen octogintaillion nine hundred forty seven novenseptuagintaillion six octoseptuagintaillion two hundred seventy two septenseptuagintaillion four hundred seventy seseptuagintaillion five hundred sixteen quinquaseptuagintaillion three hundred quattuorseptuagintaillion six hundred seventy one treseptuagintaillion eight hundred ninety seven duoseptuagintaillion nine hundred two unseptuagintaillion two hundred eighty seven septuagintaillion four hundred forty six novensexagintaillion six hundred forty four octosexagintaillion four hundred sixty one septensexagintaillion five hundred twenty two sesexagintaillion four hundred sixty three quinquasexagintaillion nine hundred forty four quattuorsexagintaillion three hundred ninety one tresexagintaillion six hundred thirty eight duosexagintaillion two hundred seventy unsexagintaillion four hundred eighty six sexagintaillion five hundred ninety seven novenquinquagintaillion one hundred sixty eight octoquinquagintaillion six hundred twenty two septenquinquagintaillion five hundred eighty sequinquagintaillion four hundred fifty six quinquaquinquagintaillion four hundred sixty quattuorquinquagintaillion one hundred eighty four tresquinquagintaillion three hundred forty six duoquinquagintaillion three hundred sixty two unquinquagintaillion four hundred thirty seven quinquagintaillion three hundred ninety two novenquadragintaillion five hundred twenty octoquadragintaillion nine hundred seventy five septenquadragintaillion three hundred sixty two sequadragintaillion eight hundred fifty four quinquaquadragintaillion seven hundred fifty six quattuorquadragintaillion two hundred ninety three tresquadragintaillion seven hundred eight duoquadragintaillion five hundred thirty four unquadragintaillion nine hundred thirty one quadragintaillion five hundred fifty five noventrigintaillion four hundred nineteen octotrigintaillion four hundred thirty eight septentrigintaillion ninety six setrigintaillion seven hundred twenty six quinquatrigintaillion six hundred twenty two quattuortrigintaillion five hundred thirty eight trestrigintaillion nine hundred thirty one duotrigintaillion nine hundred ninety three untrigintaillion two hundred ninety nine trigintaillion six hundred seventy two novemvigintiillion five hundred seventy three octovigintiillion six hundred thirty one septemvigintiillion five hundred twelve sevigintiillion seven hundred seventy four quinquavigintiillion three hundred fifty nine quattuorvigintiillion five hundred seventy four tresvigintiillion five hundred ninety one duovigintiillion six hundred seventy seven unvigintiillion seven hundred twelve vigintillion two hundred eighty eight novemdecillion one hundred forty five octodecillion eight hundred thirteen septendecillion eight sexdecillion three hundred thirty eight quindecillion five hundred sixty quattuordecillion one hundred fourteen tredecillion two hundred forty four duodecillion seven hundred eighty four undecillion eight hundred ten decillion one hundred ninety two nonillion six hundred twenty two octillion six hundred seventy septillion three hundred sixty seven sextillion five hundred twenty six quintillion five hundred ninety quadrillion one hundred twenty four trillion nine hundred forty seven billion one hundred eighty eight million four hundred seventy thousand seven hundred thirty seven
}
