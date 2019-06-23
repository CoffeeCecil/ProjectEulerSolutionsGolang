package primes

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"
	"runtime"
	"sync"
)

const E_CONST = "2.718281828459045235360287471352662497757247093699959574966967627724076630353547594571382178525166427427466391932003059921817413596629043572900334295260595630738132328627943490763233829880753195251019011573834187930702154089149934884167509244761460668082264800168477411853742345442437107539077744992069551702761838606261331384583000752044933826560297606737113200709328709127443747047230696977209310141692836819025515108657463772111252389784425056953696770785449969967946864454905987931636889230098793127736178215424999229576351482208269895193668033182528869398496465105820939239829488793320362509443117301238197068416140397019837679320683282376464804295311802328782509819455815301756717361332069811250996181881593041690351598888519345807273866738589422879228499892086805825749279610484198444363463244968487560233624827041978623209002160990235304369941849146314093431738143640546253152096183690888707016768396424378140592714563549061303107208510383750510115747704171898610687396965521267154688957035035402123407849819334321068170121005627880235193033224745015853904730419957777093503660416997329725088687696640355570716226844716256079882651787134195124665201030592123667719432527867539855894489697096409754591856956380236370162112047742722836489613422516445078182442352948636372141740238893441247963574370263755294448337998016125492278509257782562092622648326277933386566481627725164019105900491644998289315056604725802778631864155195653244258698294695930801915298721172556347546396447910145904090586298496791287406870504895858671747985466775757320568128845920541334053922000113786300945560688166740016984205580403363795376452030402432256613527836951177883863874439662532249850654995886234281899707733276171783928034946501434558897071942586398772754710962953741521115136835062752602326484728703920764310059584116612054529703023647254929666938115137322753645098889031360205724817658511806303644281231496550704751025446501172721155519486685080036853228183152196003735625279449515828418829478761085263981395599006737648292244375287184624578036192981971399147564488262603903381441823262515097482798777996437308997038886778227138360577297882412561190717663946507063304527954661855096666185664709711344474016070462621568071748187784437143698821855967095910259686200235371858874856965220005031173439207321139080329363447972735595527734907178379342163701205005451326383544000186323991490705479778056697853358048966906295119432473099587655236812859041383241160722602998330535370876138939639177957454016137223618789365260538155841587186925538606164779834025435128439612946035291332594279490433729908573158029095863138268329147711639633709240031689458636060645845925126994655724839186564209752685082307544254599376917041977780085362730941710163434907696423722294352366125572508814779223151974778060569672538017180776360346245927877846585065605078084421152969752189087401966090665180351650179250461950136658543663271254963990854914420001457476081930221206602433009641270489439039717719518069908699860663658323227870937650226014929101151717763594460202324930028040186772391028809786660565118326004368850881715723866984224220102495055188169480322100251542649463981287367765892768816359831247788652014117411091360116499507662907794364600585194199856016264790761532103872755712699251827568798930276176114616254935649590379804583818232336861201624373656984670378585330527583333793990752166069238053369887956513728559388349989470741618155012539706464817194670834819721448889879067650379590366967249499254527903372963616265897603949857674139735944102374432970935547798262961459144293645142861715858733974679189757121195618738578364475844842355558105002561149239151889309946342841393608038309166281881150371528496705974162562823609216807515017772538740256425347087908913729172282861151591568372524163077225440633787593105982676094420326192428531701878177296023541306067213604600038966109364709514141718577701418060644363681546444005331608778314317444081194942297559931401188868331483280270655383300469329011574414756313999722170380461709289457909627166226074071874997535921275608441473782330327033016823719364800217328573493594756433412994302485023573221459784328264142168487872167336701061509424345698440187331281010794512722373788612605816566805371439612788873252737389039289050686532413806279602593038772769778379286840932536588073398845721874602100531148335132385004782716937621800490479559795929059165547050577751430817511269898518840871856402603530558373783242292418562564425502267215598027401261797192804713960068916382866527700975276706977703643926022437284184088325184877047263844037953016690546593746161932384036389313136432713768884102681121989127522305625675625470172508634976536728860596675274086862740791285657699631378975303466061666980421826772456053066077389962421834085988207186468262321508028828635974683965435885668550377313129658797581050121491620765676995065971534476347032085321560367482860837865680307306265763346977429563464371670939719306087696349532884683361303882943104080029687386911706666614680001512114344225602387447432525076938707777519329994213727721125884360871583483562696166198057252661220679754062106208064988291845439530152998209250300549825704339055357016865312052649561485724925738620691740369521353373253166634546658859728665945113644137033139367211856955395210845840724432383558606310680696492485123263269951460359603729725319836842336390463213671011619282171115028280160448805880238203198149309636959673583274202498824568494127386056649135252670604623445054922758115170931492187959271800194096886698683703730220047531433818109270803001720593553052070070607223399946399057131158709963577735902719628506114651483752620956534671329002599439766311454590268589897911583709341937044115512192011716488056694593813118384376562062784631049034629395002945834116482411496975832601180073169943739350696629571241027323913874175492307186245454322203955273529524024590380574450289224688628533654221381572213116328811205214648980518009202471939171055539011394331668151582884368760696110250517100739276238555338627255353883096067164466237092264680967125406186950214317621166814009759528149390722260111268115310838731761732323526360583817315103459573653822353499293582283685100781088463434998351840445170427018938199424341009057537625776757111809008816418331920196262341628816652137471732547772778348877436651882875215668571950637193656539038944936642176400312152787022236646363575550356557694888654950027085392361710550213114741374410613444554419210133617299628569489919336918472947858072915608851039678195942983318648075608367955149663644896559294818785178403877332624705194505041984774201418394773120281588684570729054405751060128525805659470304683634459265255213700806875200959345360731622611872817392807462309468536782310609792159936001994623799343421068781349734695924646975250624695861690917857397659519939299399556754271465491045686070209901260681870498417807917392407194599632306025470790177452751318680998228473086076653686685551646770291133682756310722334672611370549079536583453863719623585631261838715677411873852772292259474337378569553845624680101390572787101651296663676445187246565373040244368414081448873295784734849000301947788802046032466084287535184836495919508288832320652212810419044804724794929134228495197002260131043006241071797150279343326340799596053144605323048852897291765987601666781193793237245385720960758227717848336161358261289622611812945592746276713779448758675365754486140761193112595851265575973457301533364263076798544338576171533346232527057200530398828949903425956623297578248873502925916682589445689465599265845476269452878051650172067478541788798227680653665064191097343452887833862172615626958265447820567298775642632532159429441803994321700009054265076309558846589517170914760743713689331946909098190450129030709956622662030318264936573369841955577696378762491885286568660760056602560544571133728684020557441603083705231224258722343885412317948138855007568938112493538631863528708379984569261998179452336408742959118074745341955142035172618420084550917084568236820089773945584267921427347756087964427920270831215015640634134161716644806981548376449157390012121704154787259199894382536495051477137939914720521952907939613762110723849429061635760459623125350606853765142311534966568371511660422079639446662116325515772907097847315627827759878813649195125748332879377157145909106484164267830994972367442017586226940215940792448054125536043131799269673915754241929660731239376354213923061787675395871143610408940996608947141834069836299367536262154524729846421375289107988438130609555262272083751862983706678722443019579379378607210725427728907173285487437435578196651171661833088112912024520404868220007234403502544820283425418788465360259150644527165770004452109773558589762265548494162171498953238342160011406295071849042778925855274303522139683567901807640604213830730877446017084268827226117718084266433365178000217190344923426426629226145600433738386833555534345300426481847398921562708609565062934040526494324426144566592129122564889356965500915430642613425266847259491431423939884543248632746184284665598533231221046625989014171210344608427161661900125719587079321756969854401339762209674945418540711844643394699016269835160784892451405894094639526780735457970030705116368251948770118976400282764841416058720618418529718915401968825328930914966534575357142731848201638464483249903788606900807270932767312758196656394114896171683298045513972950668760474091542042842999354102582911350224169076943166857424252250902693903481485645130306992519959043638402842926741257342244776558417788617173726546208549829449894678735092958165263207225899236876845701782303809656788311228930580914057261086588484587310165815116753332767488701482916741970151255978257270740643180860142814902414678047232759768426963393577354293018673943971638861176420900406866339885684168100387238921448317607011668450388721236436704331409115573328018297798873659091665961240202177855885487617616198937079438005666336488436508914480557103976521469602766258359905198704230017946553679"

const LN10CNST string = "2.30258509299404568401799145468436420760110148862877297603332790096757260967735248023599"

const LN2CONST string = "0.693147180559945309417232121458176568075500134360255254120680009493393621969694715605863326996418687"

const PICONST string = "3.14159265358979323846264338327950288419716939937510582097494459230781640628620899862803482534211706798214"

func TwoBf() func(uint) *big.Float {
	re := big.NewFloat(2)
	return func(prec uint) *big.Float {
		re.SetPrec(prec)
		return re
	}
}

func OneBf() func(uint) *big.Float {
	re := big.NewFloat(1)
	return func(prec uint) *big.Float {
		re.SetPrec(prec)
		return re
	}
}

func ZeroBf() func(uint) *big.Float {
	re := big.NewFloat(0)
	return func(prec uint) *big.Float {
		re.SetPrec(prec)
		return re
	}
}

func MakeArbitaryPrecConst(init_prec uint, s string) func(uint) *big.Float {
	rc := big.NewRat(1, 1)
	rc.SetString(s[0:init_prec])
	re := big.NewFloat(0)
	re_set := false
	prev_prec := uint(0)
	return func(prec uint) *big.Float {
		if (prev_prec == prec) && re_set {
			return re
		}
		re.SetPrec(prec)
		re.SetRat(rc)
		re_set = true
		prev_prec = prec
		return re
	}
}

func Ln10Cf(init_prec uint) func(uint) *big.Float {
	return MakeArbitaryPrecConst(init_prec, LN10CNST)
}

var ln10f = Ln10Cf(50)

func Ln2Cf(init_prec uint) func(uint) *big.Float {
	return MakeArbitaryPrecConst(init_prec, LN2CONST)
}

var ln2f = Ln2Cf(50)

func EConstf(init_prec uint) func(uint) *big.Float {
	rc := big.NewRat(1, 1)
	rc.SetString(E_CONST[0:init_prec])
	re := big.NewFloat(0)
	re_set := false
	prev_prec := uint(0)
	return func(prec uint) *big.Float {
		if prev_prec == prec && re_set {
			return re
		}
		re.SetPrec(prec)
		re.SetRat(rc)
		re_set = true
		prev_prec = prec
		return re
	}
}

func OneBi() func() *big.Int {
	re := big.NewInt(1)
	return func() *big.Int {
		return re
	}
}

var onei = OneBi()

var zerof func(uint) *big.Float = ZeroBf()
var onef func(uint) *big.Float = OneBf()
var twof func(uint) *big.Float = TwoBf()

var constef func(uint) *big.Float = EConstf(200)

//returns a channel that computes up to int 32 according to wikipedia (actually more)
func toMiller32() chan *big.Int {
	_ = fmt.Println
	ch := make(chan *big.Int,8)
	go func() {
		ch <- big.NewInt(2)
		ch <- big.NewInt(3)
		ch <- big.NewInt(5)
		ch <- big.NewInt(7)
		ch <- big.NewInt(11)
		close(ch)
	}()
	return ch
}

//returns a channel that computes up to int 64 according to wikipedia (actually, more.)
func toMiller64() chan *big.Int {
	//channel needs to have more syncronous threads
	//than there are numbers or it will deadlock.
	ch := make(chan *big.Int, 16)
	func() {
		ch <- big.NewInt(2)
		ch <- big.NewInt(3)
		ch <- big.NewInt(5)
		ch <- big.NewInt(7)
		ch <- big.NewInt(11)
		ch <- big.NewInt(13)
		ch <- big.NewInt(17)
		ch <- big.NewInt(19)
		ch <- big.NewInt(23)
		ch <- big.NewInt(29)
		ch <- big.NewInt(31)
		ch <- big.NewInt(37)
		close(ch)
	}()
	return ch
}

//MillerRabin642 is an up to 64 bit test, that does not use channels.
func MillerRabin642(n *big.Int) bool{
	dat := [...] int64{2 ,3, 5 ,7, 11, 13, 17, 19, 23, 29, 31, 37}
	r, d, t := big.NewInt(0), big.NewInt(1), big.NewInt(0)
	tmod := big.NewInt(0)
	zero := big.NewInt(0)
	one := big.NewInt(1)

	two := big.NewInt(2)
	i := big.NewInt(0)
	//check if n is odd
	//if it isn't, terminate the function.
	if n.Cmp(two) < 0 {
		return false
	}
	if n.Cmp(big.NewInt(4)) < 0 {
		return true
	}
	
	//this is safe to use because any negative number is going
	//to be returned as 'not prime' anyways.
	if n.Bit(0) == 0 {
		return false
	}

	//make t = n -1, which will be even.
	//fmt.Print
	t.Sub(n, one)
	//tmod is set to 1, so set it to be zero
	// while 2 divides t, t<- t/2, r +=1
	for tmod.Set(zero); tmod.Cmp(zero) == 0; tmod.Mod(t, two) {
		//fmt.Print("*")
		t.Div(t, two)
		r.Add(r, one)
	}
	tmod.Sub(n, one)
	i.Exp(two, r, nil)
	d.Div(tmod, i)
	halt := big.NewInt(0)
	halt.Sub(n, two)
	continueloop := false
	//halt2 := big.newFloat(n)
	x:=big.NewInt(0)
	for j, ex := 0, len(dat) ; j < ex; j++ {
		x.SetInt64( dat[j] )
		if x.Cmp(n) == 0 {
			return true
		}
		continueloop = false
		t.Exp(x, d, n)
		if (t.Cmp(one) == 0) || t.Cmp(tmod) == 0 {
			continue
		}
		for i.Set(one); i.Cmp(r) < 0; i.Add(i, one) {
			t.Exp(t, two, n)
			if t.Cmp(tmod) == 0 {
				continueloop = true
				break
			}
		}
		if continueloop{
			continue
		}
		return false
	}
	return true
}

//MillerRabin644 is an up to 64 bit test, that does not use channels.
func MillerRabin644(n *big.Int) bool{
	dat := [...] int64{2 ,3, 5 ,7, 11, 13, 17, 19, 23, 29, 31, 37}
	r, d, t := big.NewInt(0), big.NewInt(1), big.NewInt(0)
	tmod := big.NewInt(0)
	zero := big.NewInt(0)
	one := big.NewInt(1)
	var mutex = &sync.Mutex{}
	two := big.NewInt(2)
	i := big.NewInt(0)
	//check if n is odd
	//if it isn't, terminate the function.
	if n.Cmp(two) < 0 {
		return false
	}
	if n.Cmp(big.NewInt(4)) < 0 {
		return true
	}
	
	//this is safe to use because any negative number is going
	//to be returned as 'not prime' anyways.
	if n.Bit(0) == 0 {
		return false
	}

	//make t = n -1, which will be even.
	//fmt.Print
	t.Sub(n, one)
	//tmod is set to 1, so set it to be zero
	// while 2 divides t, t<- t/2, r +=1
	for tmod.Set(zero); tmod.Cmp(zero) == 0; tmod.Mod(t, two) {
		//fmt.Print("*")
		t.Div(t, two)
		r.Add(r, one)
	}
	tmod.Sub(n, one)
	i.Exp(two, r, nil)
	d.Div(tmod, i)
	halt := big.NewInt(0)
	halt.Sub(n, two)
	continueloop := false
	//halt2 := big.newFloat(n)
	ch := make(chan(byte), 8)
	lockf := true
	for j, ex := 0, len(dat) ; j < ex; j++ {
		if lockf{
			go func(ji int){
				if !lockf {
					return
				}
				x:=big.NewInt(0)
				x.SetInt64( dat[ji] )
				if x.Cmp(n) == 0 {
					mutex.Lock()
					if lockf{
						ch<-1 
					}
					mutex.Unlock()
					return
				}
				continueloop = false
				temp:= big.NewInt(0)
				temp.Exp(x, d, n)
				if (temp.Cmp(one) == 0) || temp.Cmp(tmod) == 0 {
					mutex.Lock()
					if lockf{
						ch <-2
					}
					mutex.Unlock()
					return
				}
				if !lockf {
					return
				}
				for i.Set(one); i.Cmp(r) < 0; i.Add(i, one) {
					temp.Exp(t, two, n)
					if temp.Cmp(tmod) == 0 {
						continueloop = true
						break
					}
				}

				if continueloop{
					mutex.Lock()
					if lockf{
						ch <- 2
					}
					mutex.Unlock()
				}
				
				mutex.Lock()
				if lockf{
					ch <- 0
				}
				mutex.Unlock()
			}(j)
		}
	}
	b:=byte(4)
	for i:= 0; i < len(dat) ; {
		select {
		case b = <-ch :
				if b==2{
					b=4	
					i++
					continue
				}
				if b==1{
					lockf = false
					return true
				}
				if b==0{
					lockf = false
					return false
				}
		}

	}
	return true
}


//RandomBigSequence makes an array of random big integers that are not equal to one-another
//this code is not gauranteed to terminate in a reasonable manner because it is assigning arbitrarily sized integers to a set. Have fun.
func RandomBigSequence(n int, maxsize *big.Int) chan *big.Int {
	ch := make(chan *big.Int)
	go func() {
		rng := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
		i := 0
		arr := make([]*big.Int, n)
		for i < n {

			//make a new big int pointer.
			z := big.NewInt(0)
			z.Rand(rng, maxsize)
			//set check
			for _, v := range arr {
				if z.Cmp(v) == 0 {
					continue
				}
			}
			//assign and iterate
			ch <- z
			arr[i] = z
			i++
		}
	}()
	close(ch)
	return ch
}

//MillerRabin64 Performs MillerRabin test over reasonable prime values.
//Expect all 64bit values to be verifiable
func MillerRabin64(n *big.Int) bool{
	return Miller_Rabin_Custom(n, toMiller64(), runtime.NumCPU())
}

//MillerRabin643 Performs MillerRabin test over reasonable prime values.
//Expect all 64bit values to be verifiable
func MillerRabin643(n *big.Int) bool{
	return MillerRabinCustom2(n, toMiller64())
}

//MillerRabin32 Performs MillerRabin test over reasonable prime values.
//Expect all 32bit values to be verifiable.
func MillerRabin32(n *big.Int) bool{
	return Miller_Rabin_Custom(n, toMiller32(), runtime.NumCPU())
}

//Miller_Rabin_Custom, takes an integer and a channel of numbers that it will perform
//the miller rabin test on to determine if the number is composite.
/*takes an integer n and a function returning a
//channel.
* returns true if it cannot determine primality, false if the number is composite.
*/
//Inputs a big integer and a channel to perform miller rabin testing over.
//Outputs: whether or not a number is prime.
func Miller_Rabin_Custom(n *big.Int, ch chan *big.Int, maxproc int) bool {
	runtime.GOMAXPROCS(maxproc)
	r, d, t := big.NewInt(0), big.NewInt(1), big.NewInt(0)
	tmod := big.NewInt(0)
	zero := big.NewInt(0)
	one := big.NewInt(1)

	two := big.NewInt(2)
	i := big.NewInt(0)
	//check if n is odd
	//if it isn't, terminate the function.
	if n.Cmp(two) < 0 {
		return false
	}
	if n.Cmp(big.NewInt(4)) < 0 {
		return true
	}
	
	//this is safe to use because any negative number is going
	//to be returned as 'not prime' anyways.
	if n.Bit(0) == 0 {
		return false
	}

	//make t = n -1, which will be even.
	//fmt.Print
	t.Sub(n, one)
	//tmod is set to 1, so set it to be zero
	// while 2 divides t, t<- t/2, r +=1
	for tmod.Set(zero); tmod.Cmp(zero) == 0; tmod.Mod(t, two) {
		//fmt.Print("*")
		t.Div(t, two)
		r.Add(r, one)
	}
	tmod.Sub(n, one)
	i.Exp(two, r, nil)
	d.Div(tmod, i)
	halt := big.NewInt(0)
	halt.Sub(n, two)
	continueloop := false
	//halt2 := big.newFloat(n)
	for x := range ch {
		if x.Cmp(n) == 0 {
			return true
		}
		continueloop = false
		t.Exp(x, d, n)
		if (t.Cmp(one) == 0) || t.Cmp(tmod) == 0 {
			continue
		}
		for i.Set(one); i.Cmp(r) < 0; i.Add(i, one) {
			t.Exp(t, two, n)
			if t.Cmp(tmod) == 0 {
				continueloop = true
				break
			}
		}
		if continueloop{
			continue
		}
		return false
	}
	return true
}


//Miller_Rabin_Custom, takes an integer and a channel of numbers that it will perform
//the miller rabin test on to determine if the number is composite.
/*takes an integer n and a function returning a
//channel.
* returns true if it cannot determine primality, false if the number is composite.
*/
//Inputs a big integer and a channel to perform miller rabin testing over.
//Outputs: whether or not a number is prime.
func MillerRabinCustom2(n *big.Int, ch chan *big.Int) bool {
	r, d, t := big.NewInt(0), big.NewInt(1), big.NewInt(0)
	tmod := big.NewInt(0)
	zero := big.NewInt(0)
	one := big.NewInt(1)

	two := big.NewInt(2)
	i := big.NewInt(0)
	//check if n is odd
	//if it isn't, terminate the function.
	if n.Cmp(two) < 0 {
		return false
	}
	if n.Cmp(big.NewInt(4)) < 0 {
		return true
	}
	
	//this is safe to use because any negative number is going
	//to be returned as 'not prime' anyways.
	if n.Bit(0) == 0 {
		return false
	}

	//make t = n -1, which will be even.
	//fmt.Print
	t.Sub(n, one)
	//tmod is set to 1, so set it to be zero
	// while 2 divides t, t<- t/2, r +=1
	for tmod.Set(zero); tmod.Cmp(zero) == 0; tmod.Mod(t, two) {
		//fmt.Print("*")
		t.Div(t, two)
		r.Add(r, one)
	}
	tmod.Sub(n, one)
	i.Exp(two, r, nil)
	d.Div(tmod, i)
	halt := big.NewInt(0)
	halt.Sub(n, two)
	continueloop := false
	notfinished := true
	//halt2 := big.newFloat(n)
	ch2 := make(chan(byte) )
	for x := range ch {
		go func(xv *big.Int){
			if !notfinished{
				return
			}
			xc :=big.NewInt(0)
			xc.Set(xv)
			if xc.Cmp(n) == 0 {
				if notfinished{
					ch2 <- 1
				}
				return
			}
			continueloop = false
			t.Exp(xc, d, n)
			if (t.Cmp(one) == 0) || t.Cmp(tmod) == 0 {
				if notfinished{
					ch2 <- 2
				}
				//equivalent to continue
				return
			}
			if !notfinished{
				return
			}

			for i.Set(one); i.Cmp(r) < 0; i.Add(i, one) {
				t.Exp(t, two, n)
				if t.Cmp(tmod) == 0 {
					continueloop = true
					if notfinished{
						break
					} else{
						return
					}
				}
			}
			if continueloop{
				if notfinished{
					ch2 <- 2
				}
			}
			if notfinished{
				ch2 <- 0
			}
		}(x)
	}

	for x := range ch {
		_ = x
		select {
		case b := <-ch2 :
			if b==2{
				continue
			} else if b==1{
				notfinished=false
				close(ch2)
				return true
			} else if b==0 {
				notfinished = false
				close(ch2)
				return false
			}
		}
	}
	return true
}

//Factorial:
//Multiply n*n-1*...*2*1
//place result in accumumulator
//for factorials less than one return 1
func Factorial(accumulator, n *big.Int) {
	one := big.NewInt(1)
	zero := big.NewInt(0)
	i := big.NewInt(0)
	i.Sub(n, one)
	accumulator.Set(n)

	if accumulator.Cmp(one) < 1 {
		accumulator.Set(one)
		return
	}
	for ; i.Cmp(zero) > 0; i.Sub(i, one) {
		accumulator.Mul(accumulator, i)
	}
}

/*
func CompositeSimpsonsLogE(x *big.Float) {

}*/

func FactorialFloat(accumulator, n *big.Float) {
	a, _ := accumulator.Int(nil)
	i, _ := n.Int(nil)
	Factorial(a, i)
	accumulator.SetInt(a)
}

//LogEUnbound
//What does it do to be unbounded: using a constant for ln10, divides until the
//value y which is between 0 and 1, where it will give predictable results.
//It counts these divisions as m
//then returns:
// ln(y) + m*ln(10)
// Which is equivalent to: ln(y*10^m), and or ln(x)
//Important to note: the precision of the constants is set to 50.
//Setting the prec higher shouldn't cause a panic the way I've done things
//but won't be effective. You've been warned!
func LogEUnbound(x *big.Float, iterations *big.Float, precision uint) *big.Float {
	l10 := ln10f(precision)
	if x.Cmp(zerof(precision)) < 1 {
		return nil
	}
	one := onef(precision)
	xred := big.NewFloat(0)
	mul10 := big.NewFloat(0)
	ten := big.NewFloat(10)
	xred.SetPrec(precision)
	mul10.SetPrec(precision)
	xred.SetPrec(precision)
	ten.SetPrec(precision)
	for xred.Set(x); xred.Cmp(twof(precision)) > 0; mul10.Add(one, mul10) {
		xred.Quo(xred, ten)
	}
	re := TaylorLogEExtended(xred, iterations, precision)
	mul10.Mul(l10, mul10)
	re.Add(mul10, re)
	return re
}

//computes the log for x, by shifting to x-1 and computing using the mercator
//series (which computes the log for y+1, for -1 < y < 1)
func MercatorLogENaive(x *big.Float, iterations *big.Float, precision uint) *big.Float {
	zero := big.NewFloat(0)
	zero.SetPrec(precision)
	one := big.NewFloat(1)
	one.SetPrec(precision)
	if x.Cmp(zero) == 0 {
		return one
	} else if x.Cmp(zero) < 0 {
		return nil
	}
	xmone := big.NewFloat(0)
	xmone.Sub(x, one)
	num := big.NewFloat(0)
	den := big.NewFloat(0)
	two := big.NewFloat(2)
	negone := big.NewFloat(-1)
	//2*(x-1)/(x+1)
	re := big.NewFloat(2)
	//part_one: 1/(2k+1)
	xmone.SetPrec(precision)
	two.SetPrec(precision)
	re.SetPrec(precision)
	num.SetPrec(precision)
	den.SetPrec(precision)
	negone.SetPrec(precision)
	num.Set(negone)
	i := big.NewFloat(1)
	i.SetPrec(precision)
	for ; i.Cmp(iterations) < 0; i.Add(one, i) {
		num.Mul(num, negone)
		num.Mul(num, xmone)
		den.Set(i)
		den.Quo(num, den)
		re.Add(re, den)
	}
	return re
}

//Taylor Series log E
func TaylorLogEExtended(x *big.Float, iterations *big.Float, precision uint) *big.Float {
	zero := big.NewFloat(0)
	zero.SetPrec(precision)
	one := onef(precision)
	if x.Cmp(zero) == 0 {
		return one
	} else if x.Cmp(zero) < 0 {
		return nil
	}
	num := big.NewFloat(0)
	den := big.NewFloat(0)
	two := twof(precision)
	//2*(x-1)/(x+1)
	re := big.NewFloat(2)
	//part_one: 1/(2k+1)
	part_one_a := big.NewFloat(1)
	//[(x-1)**2 / (x+1)**2 ]**k
	part_two_c := big.NewFloat(0)
	//since this is dependent on k, it is split into a constant and accumulated part.
	part_two_a := big.NewFloat(1)
	final_accumulator := big.NewFloat(0)
	//fmt.Println(two, two.Prec())
	re.SetPrec(precision)
	part_one_a.SetPrec(precision)
	part_two_a.SetPrec(precision)
	part_two_c.SetPrec(precision)
	final_accumulator.SetPrec(precision)
	num.SetPrec(precision)
	den.SetPrec(precision)
	i := big.NewFloat(0)
	i.SetPrec(precision)
	num.Sub(x, one)
	den.Add(x, one)
	re.Mul(re, num)
	re.Quo(re, den)
	num.Mul(num, num)
	den.Mul(den, den)
	part_two_c.Quo(num, den)
	num.Set(one)
	num.SetPrec(precision)
	for ; i.Cmp(iterations) < 0; i.Add(i, one) {
		den.Mul(two, i)
		den.Add(den, one)
		part_one_a.Quo(one, den)
		num.Mul(part_one_a, part_two_a)
		final_accumulator.Add(final_accumulator, num)
		//set part two a to the power of k by multiplying it in this final step.
		part_two_a.Mul(part_two_a, part_two_c)
	}
	re.Mul(re, final_accumulator)
	return re
}

//MakeE
//calculate e using Brothers formula.
//To some precision specified by b.
//You need precision = 5000 and p = 16*4096 to get more than 10000 accurate digits.
func MakeE(precision *big.Float, p uint) *big.Float {
	re := big.NewFloat(0)
	num := big.NewFloat(0)
	den := big.NewFloat(0)
	two := big.NewFloat(2)
	one := big.NewFloat(1)
	re.SetPrec(p)
	num.SetPrec(p)
	den.SetPrec(p)
	two.SetPrec(p)
	one.SetPrec(p)

	for i := big.NewFloat(0); i.Cmp(precision) < 0; i.Add(i, one) {
		num.Mul(two, i)
		den.Set(num)
		num.Add(two, num)
		den.Add(one, den)
		FactorialFloat(den, den)
		num.Quo(num, den)
		re.Add(re, num)
	}
	return re
}

/*
func NaturalLog(n *big.Int) *big.Int{
	rt := bit.NewInt(0)
}
*/

//Sieve of Erasthones, the classic prime finder.
func SieveEra(n int) []int {
	if n < 2 {
		return nil
	}
	if n < 3 {
		rt := make([]int, 1)
		rt[0] = 2
		return rt
	}
	if n < 4 {
		rt := make([]int, 2)
		rt[0] = 2
		rt[1] = 3
		return rt
	}

	earr, found := ErasthonesArray(n)
	rt := make([]int, found)
	rt[0] = 2
	rt[1] = 3
	j:=2
	for i := 2; j < found; i++ {
		if !earr[i] {
			rt[j] = 2*i + 1
			j++
		}
	}
	return rt
}

//ErasthonesArray Produces an array of numbers which are prime.
//Only stores the odd numbers of the array.
func ErasthonesArray(n int) ([]bool, int) {
	if n < 2 {
		return nil, -1
	}
	if n < 3 {
		rt := make([]bool, 1)
		rt[0] = true
		return rt, 1
	}
	rt := make([]bool, int(n/2)+1)
	rt[0] = true
	rt[1] = true
	if n < 4 {
		return rt, 2
	}
	rt[0] = false
	rt[1] = false
	//okay the above is confusing:
	//basically to avoid reassining an entire array of booleans, flip the values of a normal sieve of erasthones.
	k := 3
	j := 3
	f := 1
	ln := n+1
	for k < ln {
		if !rt[(k-1)/2] {
			//if k-1/2 is false then it is an odd prime.
			for j < ln {
				//is the kth odd number a prime?
				if j*k < ln {
					//if j is a prime
					//set the multiple to not prime.
					rt[(j*k-1)/2] = true
				} else {
					break
				}
				j += 2
			}
			//we found another one!
			f++
		}
		j = 3
		k += 2
	}
	return rt, f
}

//SieveNaiveEra produces an array of numbers which are prime.
func SieveNaiveEra(n int) []int {
	earr, fnd := erasthonesNaiveArray(n)
	if fnd == -1 {
		return nil
	}
	rt := make([]int, fnd)
	for i, x := 2, 0; i < n && x < fnd; i++ {
		if !earr[i] {
			rt[x] = i
			x++
		}
	}
	return rt
}

//ErasthonesNaiveArray
//returns an array. If a number is composite than it sets the number to true.
//If a number is prime then it sets the number to false
func erasthonesNaiveArray(n int) ([]bool, int) {
	if n < 2 {
		return nil, -1
	}
	rt := make([]bool, n+1)
	rt[0] = false
	rt[1] = false
	rt[2] = true
	rt[3] = true
	if n < 3 {
		return rt[:2], 1
	}
	if n < 4 {
		return rt, 2
	}
	rt[0] = true
	rt[1] = true
	rt[2] = false
	rt[3] = false
	f := 0
	ln:= n+1
	for i := 2; i < ln; i++ {
		if !rt[i] {
			for j := 2; j < n; j++ {
				if i*j < ln {
					rt[i*j] = true
				} else {
					break
				}
			}
			f++
		}
	}
	return rt, f
}

//Returns an array of integers which are prime, up to a big.Integer limit.
//Returns based off the nth prime that they are.
func OddballPrimesTest(limit int, sizecutoff *big.Int )[]bool {
	//two:= big.NewInt(2)
	//make sure that the limit is set to the max size	
	l := limit
	if l < 0 && sizecutoff == nil {
		return nil
	}
	if l == 0 {
		cl := big.NewInt(0)
		cl.Rsh(sizecutoff, 1)
		if cl.IsInt64(){
			if cl.Int64() > int64(int(^uint(0) >> 1)){
				l = int(^uint(0) >> 1)
			}else{
				l = int(cl.Int64())
			}
		} else {
			l = int(^uint(0) >> 1)
		}
	}
	numbers:= make([]bool, l)
	zero:= big.NewInt(0)
	cnt3 := 0
	cnt5 := 0
	cnt7 := 0
	cnt11 := 0
	if l < 1 {
		return nil
	}

	numbers[0]=true//two is prime
	if l > 0{
		numbers[1] = true
	} else{
		return numbers
	}
	if l > 1{
		numbers[2] = true
	} else{
		return numbers
	}
	if l > 2{
		numbers[3] = true
	} else{
		return numbers
	}
	if l > 4{
		numbers[5] = true
	} else{
		return numbers
	}
	if sizecutoff.Cmp(big.NewInt(11)) < 0 {
		return numbers
	}
	cnt3 = 2
	cnt5 = 4
	cnt7 = 3
	cnt11 = 1
	iv := big.NewInt(0)
	i:=6
	setiv := func(){
		iv.SetInt64(int64(i))
		iv.Lsh(iv, 1)
		iv.Add(iv, onei())
	}
	notcomposite:= true
	testAndSet:= func(n *int, max int){
		if *n == max{
			numbers[i] = false
			notcomposite = false
			*n = 0
		}
	}
	pmod := big.NewInt(0)	
	j:=1
	jv:=big.NewInt(0)
	setjv := func(){
		jv.SetInt64(int64(j))
		jv.Lsh(jv, 1)
		jv.Add(jv, onei())
	}
	
	innerloop := func(){
		numbers[i] = true
		//Assume the number is prime.
		//from j=1 to i, set jv=2*j + 1
		//if that number is not composite.
		for j=1; j < i; j++{
			if(numbers[j]){
				setjv()
				pmod.Mod(iv,jv)
				if pmod.Cmp(zero)==0{
					numbers[i] = false
					return
				}
			}
		}
	}

	for setiv() ;  ; setiv() {
		if (sizecutoff != nil){
			if iv.Cmp(sizecutoff) > 0 {
				return numbers
			}
		}
		if (limit != 0 ){
			if i > limit {
				return numbers
			}
		}
		testAndSet(&cnt3, 3)//if cnt3, then the number is divisible by 3
		testAndSet(&cnt5,5)//if cnt5, then the number is divisible by 5
		testAndSet(&cnt7,7)//ect.
		testAndSet(&cnt11,11)
		if notcomposite{
			innerloop()
		}
		cnt3++
		cnt5++
		cnt7++
		cnt11++
		i++
		notcomposite = true
	}
	//return numbers
}

func perfectPowersTester(n *big.Int) bool {
	//i := big.NewInt(0)
	//i.Sqrt(n)
	t := big.NewInt(0)
	//t.Mul(i,i)
	//check to see if t is the square root of two. If not, you know something real important.
	//Observe sqrt(x) => x**(n/2)
	/*
	If there is some number i such that i**b == n
	and b is even then it can be rewritten such that x**(n/2) = i**b
	So we only need to check odd numbers under i once we've computed this.
	*/
	/*if( t.Cmp(n) == 0){
		return true
	}*/
	OneI := big.NewInt(1)
	t.SetInt64(1)
	k := big.NewInt(0)//set k to the number of bits in l.
	k.Sqrt(n)
	//this is the ceiling of log_2(n)
	lowsearch := big.NewInt(0)
	//highsearch := big.NewInt(0)
	stopsearch := big.NewInt(0)
	//beam := big.NewInt(0)
	//Tried binary search per https://ask.sagemath.org/question/44275/implementing-the-aks-primality-test/
	//It failed.
	/*
	What definitely works:
	Doing a linear search from 2 .. sqrt(n)
	and seeing if x**n = some perfect power.
	*/
	for ki := big.NewInt(2); ki.Cmp(k) < 1 ; ki.Add(ki, OneI) {
		lowsearch.SetInt64(0)
		for stopsearch.SetInt64(int64(n.BitLen())); lowsearch.Cmp(stopsearch) < 1; lowsearch.Add(lowsearch, OneI) {
			t.Exp(ki, lowsearch, nil )
			if t.Cmp(n) == 0{
				return true
			} else if t.Cmp(n) > 0 {
				continue
			}
			
		}
	}
	return false
}

//PerfectPowerAKSCheck checks is a number is a perfect power of an odd number, or is divisble by two.
func PerfectPowerAKSCheck(n *big.Int) bool {
		//i := big.NewInt(0)
	//i.Sqrt(n)
	t := big.NewInt(0)
	//t.Mul(i,i)
	//check to see if t is the square root of two. If not, you know something real important.
	//Observe sqrt(x) => x**(n/2)
	/*
	If there is some number i such that i**b == n
	and b is even then it can be rewritten such that x**(n/2) = i**b
	So we only need to check odd numbers under i once we've computed this.
	*/
	/*if( t.Cmp(n) == 0){
		return true
	}*/
	OneI := big.NewInt(1)
	TwoI := big.NewInt(2)
	k := big.NewInt(0)//set k to the number of bits in l.
	k.Sqrt(n)
	//quickly check if nroot of t is a square root.
	t.Mul(k,k)
	if(t.Cmp(n) == 0){
		return true
	}
	//this is the ceiling of log_2(n)
	lowsearch := big.NewInt(0)
	//highsearch := big.NewInt(0)
	stopsearch := big.NewInt(0)
	//beam := big.NewInt(0)
	//Tried binary search per https://ask.sagemath.org/question/44275/implementing-the-aks-primality-test/
	//It failed.
	/*
	What definitely works:
	Doing a linear search with i E ( 2 .. sqrt(n) )
	and seeing if x**i = some perfect power.
	*/
	for ki := big.NewInt(2); ki.Cmp(k) < 1 ; ki.Add(ki, OneI) {
		lowsearch.SetInt64(1)//since we already checked the square root we can set this to 3.
		for stopsearch.SetInt64(int64(n.BitLen())); lowsearch.Cmp(stopsearch) < 1; lowsearch.Add(lowsearch, TwoI) {
			//t = QuickExp(ki, lowsearch) //too slow in testing.
			t.Exp(ki, lowsearch, nil )
			if t.Cmp(n) == 0{
				return true
			} else if t.Cmp(n) > 0 {
				continue
			}
		}
	}
	return false
}

////Too slow
// func QuickExp(n *big.Int, i *big.Int) *big.Int{
// 	/*
// 	Comput n**i
// 	*/
// 	oe := big.NewInt(0)
// 	oe.Set(i)//set the original value to n
// 	t:= big.NewInt(0)
// 	t.Set(n)
// 	rt := big.NewInt(1)
// 	OneI := big.NewInt(1)
// 	for oe.Cmp(OneI) > 0 {
// 		//check to see if the ov is odd.
// 		if ( oe.Bit(0) == 1 ){
// 			rt.Mul(rt, t)
// 		}
// 		t.Mul(t,t)
// 		oe.Rsh(oe, 1)
// 	}
// 	//Original exponent is either 1, or 0. If a is not 0:
// 	if oe.Bit(0) == 1 {
// 		rt.Mul(rt, t)
// 	}
//     return rt
// }

func MultiplicativeOrder(n , a *big.Int) *big.Int{
	t := big.NewInt(1)
	rt := big.NewInt(1)
	OneI:=big.NewInt(1)
	for ; rt.Cmp(n) < 0 ; rt.Add(rt, OneI){
		t := t.Exp(a, rt, n)
		if t.Cmp(OneI) == 0{
			return rt
		}
	}
	rt.Sub(n, OneI)
	return rt
}
/*
func AKSPrimeTest(n *big.Int) bool {
	//if n % 2 == 1 its first bit will be set to 1.
	if n.Bit(0) == 0 {
		return false
	}

	//check if n is a perfect power.
	if PerfectPowerAKSCheck(n) {
		return false
	}
	//Compute r such that ord_r (n) > log_2(n)**2. Note,
	//this uses the gcd function to ensure that gcd(n,t) == 1,
	//which may be inefficient. Another attempt will try with a naked order function.
	
	OneI := big.NewInt(1)
	rt :=big.NewInt(0)
	s := big.NewInt(0)
	//Set u to be rounded log2 value of n.
	//Note that log2(2) == 1, log2(1) == 0, and this code honors that.
	u := big.NewInt( int64(n.BitLen()) )
	u.Sub(u, OneI)
	u.Mul(u,u)
	
	for t :=big.NewInt(2) ; t.Cmp(n) < 0 ; t.Add(t, OneI){
		s.GCD(nil, nil, t, n)
		if (s.Cmp(OneI) != 0){
			return false
		}
		rt = MultiplicativeOrder(n, t)
		if (rt.Cmp(u) > 1){
			break
		}
	}
	//Check that for 2 to r, n-1 no number divides n.
	u = u.Sub(n, OneI)
	if(rt.Cmp(u) == 1){
		u.Set(rt)
	}
	w:=big.NewInt(0)
	ZeroI := big.NewInt(0)
	for t:= big.NewInt(2); t.Cmp(u) < 1; t.Add(t, OneI){
		s.DivMod(u, t, w)
		//if there is no residue, it cannot be prime.
		//since a number divided it.
		if(w.Cmp(ZeroI) == 0){
			return false
		}
	}
	//if rt >= n, it is a prime.
	if rt.Cmp(n) > 0{
		return true
	}
	//Setup to check polynomial expansion.
	u.Sqrt(rt)
	w.SetInt64( int64(n.BitLen()))
	w.Sub(w, OneI)
	u.Mul(u,w)
	//set to sqrtn * log n. Note that literature suggests this is
	//floor ( sqrt(n) * log n ), the values gotten by these functions will be
	//floor(sqrt(n)* floog(log_2(n))
	for t:=big.NewInt(1) ; t.Cmp(u) < 0; t.Add(t, OneI){
		poly := polycoef.Polynomial(n.Int64())
		for _, c := range poly {
			
		}
	}
	return true
}
*/
