package stdlib

import (
	"fmt"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"
)

var strs = []string{
	"In 629 AD, during the Battle of Mu'tah in what is today Al-Karak, the Byzantines and their Arab Christian clients, the Ghassanids, staved off an attack by a Muslim Rashidun force that marched northwards towards the Levant from the Hejaz (in modern-day Saudi Arabia).[51] The Byzantines however were defeated by the Muslims in 636 AD at the decisive Battle of Yarmouk just north of Transjordan.[51] Transjordan was an essential territory for the conquest of Damascus.[52] The first, or Rashidun, caliphate was followed by that of the Ummayads (661–750).[52] Under the Umayyad Caliphate, several desert castles were constructed in Transjordan, including: Qasr Al-Mshatta and Qasr Al-Hallabat.[52] The Abbasid Caliphate's campaign to take over the Umayyad's began in Transjordan.[53] A powerful 747 AD earthquake is thought to have contributed to the Umayyads defeat to the Abbasids, who moved the caliphate's capital from Damascus to Baghdad.[53] During Abbasid rule (750–969), several Arab tribes moved northwards and settled in the Levant.[52] Concurrently, growth of maritime trade diminished Transjordan's central position, and the area became increasingly impoverished.[54] After the decline of the Abbasids, Transjordan was ruled by the Fatimid Caliphate (969–1070), then by the Crusader Kingdom of Jerusalem (1115–1187).[55]",
	"The Karak Castle (c. 12th century AD) built by the Crusaders, and later expanded under the Muslim Ayyubids and Mamluks.",
	"The Crusaders constructed several Crusader castles as part of the Lordship of Oultrejordain, including those of Montreal and Al-Karak.[56] The Ayyubids built the Ajloun Castle and rebuilt older castles, to be used as military outposts against the Crusaders.[57] During the Battle of Hattin (1187) near Lake Tiberias just north of Transjordan, the Crusaders lost to Saladin, the founder of the Ayyubid dynasty (1187–1260).[57] Villages in Transjordan under the Ayyubids became important stops for Muslim pilgrims going to Mecca who travelled along the route that connected Syria to the Hejaz.[58] Several of the Ayyubid castles were used and expanded by the Mamluks (1260–1516), who divided Transjordan between the provinces of Karak and Damascus.[59] During the next century Transjordan experienced Mongol attacks, but the Mongols were ultimately repelled by the Mamluks after the Battle of Ain Jalut (1260).[60]",
	"In 1516, the Ottoman Caliphate's forces conquered Mamluk territory.[61] Agricultural villages in Transjordan witnessed a period of relative prosperity in the 16th century, but were later abandoned.[62] Transjordan was of marginal importance to the Ottoman authorities.[63] As a result, Ottoman presence was virtually absent and reduced to annual tax collection visits.[62] More Arab bedouin tribes moved into Transjordan from Syria and the Hejaz during the first three centuries of Ottoman rule, including the Adwan, the Bani Sakhr and the Howeitat.[64] These tribes laid claims to different parts of the region, and with the absence of a meaningful Ottoman authority, Transjordan slid into a state of anarchy that continued till the 19th century.[65] This led to a short-lived occupation by the Wahhabi forces (1803–1812), an ultra-orthodox Islamic movement that emerged in Najd (in modern-day Saudi Arabia).[66] Ibrahim Pasha, son of the governor of the Egypt Eyalet under the request of the Ottoman sultan, rooted out the Wahhabis by 1818.[67] In 1833 Ibrahim Pasha turned on the Ottomans and established his rule over the Levant.[68] His oppressive policies led to the unsuccessful peasants' revolt in Palestine in 1834.[68] Transjordanian cities of Al-Salt and Al-Karak were destroyed by Ibrahim Pasha's forces for harbouring a peasants' revolt leader.[68] Egyptian rule was forcibly ended in 1841, with Ottoman rule restored.[68]",
	"The Ajloun Castle (c. 12th century AD) built by the Ayyubid leader Saladin for use against the Crusades.",
	"	Only after Ibrahim Pasha's campaign did the Ottoman Empire try to solidify its presence in the Syria Vilayet, which Transjordan was part of.[69] A series of tax and land reforms (Tanzimat) in 1864 brought some prosperity back to agriculture and to abandoned villages, while it provoked a backlash in other areas of Transjordan.[69] Muslim Circassians and Chechens, fleeing Russian persecution, sought refuge in the Levant.[70] In Transjordan and with Ottoman support, Circassians first settled in the long-abandoned vicinity of Amman in 1867, and later in the surrounding villages.[70] After having established its administration, conscription and heavy taxation policies by the Ottoman authorities, led to revolts in the areas it controlled.[71] Transjordan's tribes in particular revolted during the Shoubak (1905) and the Karak Revolts (1910), which were brutally suppressed.[70] The construction of the Hejaz Railway in 1908–stretching across the length of Transjordan and linking Mecca with Istanbul–helped the population economically as Transjordan became a stopover for pilgrims.[70] However, increasing policies of Turkification and centralization adopted by the Ottoman Empire disenchanted the Arabs of the Levant.[72]",
	"Soldiers of the Hashemite-led Arab Army holding the flag of the Great Arab Revolt in 1916.",
	"Four centuries of stagnation during Ottoman rule came to an end during World War I by the 1916 Arab Revolt; driven by long-term resentment towards the Ottoman authorities, and growing Arab nationalism.[70] The revolt was led by Sharif Hussein of Mecca, and his sons Abdullah, Faisal and Ali, members of the Hashemite dynasty of the Hejaz, descendants of the Prophet Muhammad.[70] Locally, the revolt garnered the support of the Transjordanian tribes, including Bedouins, Circassians and Christians.[73] The Allies of World War I, including Britain and France, whose imperial interests converged with the Arabist cause, offered support.[74] The revolt started on 5 June 1916 from Medina and pushed northwards until the fighting reached Transjordan in the Battle of Aqaba on 6 July 1917.[75] The revolt reached its climax when Faisal entered Damascus in October 1918, and established the Arab Kingdom of Syria, which Transjordan was part of.[73]",
	"The nascent Hashemite Kingdom was forced to surrender to French troops on 24 July 1920 during the Battle of Maysalun.[76] Arab aspirations failed to gain international recognition, due mainly to the secret 1916 Sykes–Picot Agreement, which divided the region into French and British spheres of influence, and the 1917 Balfour Declaration, which promised Palestine to Jews.[77] This was seen by the Hashemites and the Arabs as a betrayal of their previous agreements with the British,[78] including the 1915 McMahon–Hussein Correspondence, in which the British stated their willingness to recognize the independence of a unified Arab state stretching from Aleppo to Aden under the rule of the Hashemites.[79]:55 Abdullah, the second son of Sharif Hussein, arrived from Hejaz by train in Ma'an in southern Transjordan on 21 November 1920 to redeem the Kingdom his brother had lost.[80] Transjordan then was in disarray; widely considered to be ungovernable with its dysfunctional local governments.[81] Abdullah then moved to Amman and established the Emirate of Transjordan on 11 April 1921.[82]",
	"Al-Salt residents gather on 20 August 1920 during the British High Commissioner's visit to Transjordan.",
	"The British reluctantly accepted Abdullah as ruler of Transjordan.[83] Abdullah gained the trust of Tansjordan's tribal leaders before scrambling to convince them of the benefits of an organized government.[84] Abdullah's successes drew the envy of the British, even when it was in their interest.[85] In September 1922, the Council of the League of Nations recognised Transjordan as a state under the British Mandate for Palestine and the Transjordan memorandum, and excluded the territories east of the Jordan River from the provisions of the mandate dealing with Jewish settlement.[86][87] Transjordan remained a British mandate until 1946, but it had been granted a greater level of autonomy than the region west of the Jordan River.[88]",
	"The first organised army in Jordan was established on 22 October 1920, and was named the \"Arab Legion\".[89] The Legion grew from 150 men in 1920 to 8,000 in 1946.[90] Multiple difficulties emerged upon the assumption of power in the region by the Hashemite leadership.[89] In Transjordan, small local rebellions at Kura in 1921 and 1923 were suppressed by Emir Abdullah with the help of British forces.[89] Wahhabis from Najd regained strength and repeatedly raided the southern parts of his territory in (1922–1924), seriously threatening the Emir's position.[89] The Emir was unable to repel those raids without the aid of the local Bedouin tribes and the British, who maintained a military base with a small RAF detachment close to Amman.[89]",
	"King Abdullah I on 25 May 1946 reading the declaration of independence. The Treaty of London, signed by the British Government and the Emir of Transjordan on 22 March 1946, recognised the independence of Transjordan upon ratification by both countries' parliaments.[91] On 25 May 1946, the day that the treaty was ratified by the Transjordan parliament, Transjordan was raised to the status of a kingdom under the name of the Hashemite Kingdom of Transjordan, with Abdullah as its first king.[92] The name was shortened to the Hashemite Kingdom of Jordan on 26 April 1949.[10] Jordan became a member of the United Nations on 14 December 1955.[10]",
	"On 15 May 1948, as part of the 1948 Arab–Israeli War, Jordan invaded Palestine together with other Arab states.[93] Following the war, Jordan controlled the West Bank and on 24 April 1950 Jordan formally annexed these territories after the Jericho conference.[94][95] In response, some Arab countries demanded Jordan's expulsion from the Arab League.[94] On 12 June 1950, the Arab League declared that the annexation was a temporary, practical measure and that Jordan was holding the territory as a \"trustee\" pending a future settlement.[96] King Abdullah was assassinated at the Al-Aqsa Mosque in 1951 by a Palestinian militant, amid rumours he intended to sign a peace treaty with Israel.[97]",
	"Abdullah was succeeded by his son Talal, who would soon abdicate due to illness in favour of his eldest son Hussein.[98] Talal established the country's modern constitution in 1952.[98] Hussein ascended to the throne in 1953 at the age of 17.[97] Jordan witnessed great political uncertainty in the following period.[99] The 1950s were a period of political upheaval, as Nasserism and Pan-Arabism swept the Arab World.[99] On 1 March 1956, King Hussein Arabized the command of the Army by dismissing a number of senior British officers, an act made to remove remaining foreign influence in the country.[100] In 1958, Jordan and neighbouring Hashemite Iraq formed the Arab Federation as a response to the formation of the rival United Arab Republic between Nasser's Egypt and Syria.[101] The union lasted only six months, being dissolved after Iraqi King Faisal II (Hussein's cousin) was deposed by a bloody military coup on 14 July 1958.[101]",
	"King Hussein on 21 March 1968 checking an abandoned Israeli tank in the aftermath of the Battle of Karameh.",
	"Jordan signed a military pact with Egypt just before Israel launched a preemptive strike on Egypt to begin the Six-Day War in June 1967, where Jordan and Syria joined the war.[102] The Arab states were defeated and Jordan lost control of the West Bank to Israel.[102] The War of Attrition with Israel followed, which included the 1968 Battle of Karameh where the combined forces of the Jordanian Armed Forces and the Palestine Liberation Organization (PLO) repelled an Israeli attack on the Karameh camp on the Jordanian border with the West Bank.[102] Despite the fact that the Palestinians had limited involvement against the Israeli forces, the events at Karameh gained wide recognition and acclaim in the Arab world.[103] As a result, the time period following the battle witnessed an upsurge of support for Palestinian paramilitary elements (the fedayeen) within Jordan from other Arab countries.[103] The fedayeen activities soon became a threat to Jordan's rule of law.[103] In September 1970, the Jordanian army targeted the fedayeen and the resultant fighting led to the expulsion of Palestinian fighters from various PLO groups into Lebanon, in a conflict that became known as Black September.[103]",
	"In 1973, Egypt and Syria waged the Yom Kippur War on Israel, and fighting occurred along the 1967 Jordan River cease-fire line.[103] Jordan sent a brigade to Syria to attack Israeli units on Syrian territory but did not engage Israeli forces from Jordanian territory.[103] At the Rabat summit conference in 1974, Jordan agreed, along with the rest of the Arab League, that the PLO was the \"sole legitimate representative of the Palestinian people\".[103] Subsequently, Jordan renounced its claims to the West Bank in 1988.[103]",
	"At the 1991 Madrid Conference, Jordan agreed to negotiate a peace treaty sponsored by the US and the Soviet Union.[103] The Israel-Jordan Treaty of Peace was signed on 26 October 1994.[103] In 1997, Israeli agents entered Jordan using Canadian passports and poisoned Khaled Meshal, a senior Hamas leader.[103] Israel provided an antidote to the poison and released dozens of political prisoners, including Sheikh Ahmed Yassin after King Hussein threatened to annul the peace treaty.[103]",
	"Army Chief Habis Majali and Prime Minister Wasfi Tal during a military parade in 1970, two widely acclaimed national figures.",
	"On 7 February 1999, Abdullah II ascended the throne upon the death of his father Hussein.[104] Abdullah embarked on economic liberalisation when he assumed the throne, and his reforms led to an economic boom which continued until 2008.[105] Abdullah II has been credited with increasing foreign investment, improving public-private partnerships and providing the foundation for Aqaba's free-trade zone and Jordan's flourishing information and communication technology (ICT) sector.[105] He also set up five other special economic zones.[105] However, during the following years Jordan's economy experienced hardship as it dealt with the effects of the Great Recession and spillover from the Arab Spring.[106]",
	"Al-Qaeda under Abu Musab al-Zarqawi's leadership launched coordinated explosions in three hotel lobbies in Amman on 9 November 2005, resulting in 60 deaths and 115 injured.[107] The bombings, which targeted civilians, caused widespread outrage among Jordanians.[107] The attack is considered to be a rare event in the country, and Jordan's internal security was dramatically improved afterwards.[107] No major terrorist attacks have occurred since then.[108] Abdullah and Jordan are viewed with contempt by Islamic extremists for the country's peace treaty with Israel and its relationship with the West.[109]",
	"The Arab Spring were large-scale protests that erupted in the Arab World in 2011, demanding economic and political reforms.[110] Many of these protests tore down regimes in some Arab nations, leading to instability that ended with violent civil wars.[110] In Jordan, in response to domestic unrest, Abdullah replaced his prime minister and introduced a number of reforms including: reforming the Constitution, and laws governing public freedoms and elections.[110] Proportional representation was re-introduced to the Jordanian parliament in the 2016 general election, a move which he said would eventually lead to establishing parliamentary governments.[111] Jordan was left largely unscathed from the violence that swept the region despite an influx of 1.4 million Syrian refugees into the natural resources-lacking country and the emergence of the Islamic State of Iraq and the Levant (ISIL)",
}

// test wait group and atomic
func TestAtomicAdd(t *testing.T) {
	s := int32(1)
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			for i := 0; i < 100; i++ {
				atomic.AddInt32(&s, 1)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(s)
}

// test wait group and atomic
func TestAtomicCompareAndSwap(t *testing.T) {
	s := int32(1)
	c := make(chan int, 10)
	for i := 0; i < 10; i++ {
		go func(n int) {
			waited := 0
			for i := 0; i < 100; i++ {
				cur := atomic.LoadInt32(&s)
				for !atomic.CompareAndSwapInt32(&s, cur, cur+1) {
					cur = atomic.LoadInt32(&s)
					waited++ // count how many empty loops in each goroutine
				}
			}
			c <- waited
		}(i)
	}
	for i := 0; i < 10; i++ {
		k := <-c
		fmt.Println(k)
	}
	fmt.Println(s)
}

func TestSyncMap(t *testing.T) {
	maper := new(sync.Map)
	var wg sync.WaitGroup
	wg.Add(len(strs))
	for i := range strs {
		go func(i int) {
			tokens := strings.Split(strs[i], " ")
			for j := range tokens {
				n, ok := maper.LoadOrStore(tokens[j], 1)
				if ok {
					maper.Store(tokens[j], n.(int)+1)
				}
			}
			fmt.Println(i, "finished")
			wg.Done()
		}(i)
	}
	wg.Wait()
	maper.Range(func(key, value interface{}) bool {
		fmt.Println(key, "==", value)
		return true
	})
}

func TestSyncMap1(t *testing.T) {
	// use ordinary map, cause 'fatal error: concurrent map read and map write'
	maper := make(map[string]int)
	var wg sync.WaitGroup
	wg.Add(len(strs))
	for i := range strs {
		go func(i int) {
			tokens := strings.Split(strs[i], " ")
			for j := range tokens {
				n, ok := maper[tokens[j]]
				if ok {
					maper[tokens[j]] = n + 1
				} else {
					maper[tokens[j]] = 1
				}
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(maper)
}

func TestChanBlock(t *testing.T) {
	chan1 := make(chan rune)
	chan2 := make(chan rune)
	chan3 := make(chan rune)
	exit := make(chan rune)
	printer := func(num int, exiter bool, ch rune, chan1 <-chan rune, chan2 chan<- rune) {
		for i := 0; i < num; i++ {
			k := <-chan1
			fmt.Println(ch, k)
			if exiter && i == num-1 {
				exit <- ch
			} else {
				chan2 <- ch
			}
		}
	}
	go printer(5, false, 'a', chan1, chan2)
	go printer(5, false, 'b', chan2, chan3)
	go printer(5, true, 'c', chan3, chan1)
	chan1 <- 'm'
	fmt.Println(<-exit)
}

func TestEmptyStruct(t *testing.T) {
	// the size for empty struct is 0, and empty struct 'aa' and 'cc' are actually the same.
	// so use make(chan struct{}) rather than make(chan int) when you use channel for notification
	// other than transfer data. it can avoid allocate memory.
	aa := struct{}{}
	bb := struct{ i int }{2}
	cc := struct{}{}
	fmt.Println(unsafe.Sizeof(aa))
	fmt.Println(unsafe.Sizeof(bb))
	fmt.Println(unsafe.Sizeof(cc))
	fmt.Printf("%p\n", &aa)
	fmt.Printf("%p\n", &bb)
	fmt.Printf("%p\n", &cc)
	ch := make(chan struct{})
	go func() {
		time.Sleep(time.Second)
		fmt.Println("i exit")
		ch <- struct{}{}
	}()
	<-ch
	fmt.Println("main exit")
}
