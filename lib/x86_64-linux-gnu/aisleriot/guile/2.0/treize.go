GOOF----LE-8-2.03      ] Q 4     h      ] g  guile¤	 ¤	g  process-use-modules¤	 ¤	 ¤	g  	aisleriot¤	g  	interface¤	 ¤		 ¤	
g  api¤	
 ¤	 ¤	g  initialize-playing-area¤	g  set-ace-low¤	g  make-standard-deck¤	g  shuffle-deck¤	g  add-normal-slot¤	g  DECK¤	g  HORIZPOS¤	e  2/3¤	g  add-extended-slot¤	g  right¤	g  add-carriage-return-slot¤	g  add-blank-slot¤	g  VERTPOS¤	e  0.5¤	g  deal-cards-face-up¤										
																				 ¤	g  give-status-message¤	g  new-game¤	g  set-statusbar-message¤	 g  get-stock-no-string¤	!g  string-append¤	"g  _¤	#f  Stock left:¤	$f   ¤	%g  number->string¤	&g  length¤	'g  	get-cards¤	(g  empty-slot?¤	)g  	get-value¤	*g  king¤	+g  
available?¤	,g  button-pressed¤	-g  get-top-card¤	.g  
droppable?¤	/g  add-to-score!¤	0g  remove-card¤	1g  reverse¤	2g  
set-cards!¤	3g  button-released¤	4 ¤	5	 ¤	6g  button-clicked¤	7g  button-double-clicked¤	8g  game-won¤	9g  get-hint¤	:g  game-continuable¤	;g  club¤	<f  Remove the king of clubs.¤	=g  diamond¤	>f  Remove the king of diamonds.¤	?g  heart¤	@f  Remove the king of hearts.¤	Ag  spade¤	Bf  Remove the king of spades.¤	Cg  hint-remove-king¤	Dg  
check-move¤	Eg  
hint-click¤	Fg  get-suit¤	Gg  	hint-move¤	Hf  Deal a card¤	Ig  	dealable?¤	Jg  check-waste¤	Kg  get-options¤	Lg  apply-options¤	Mg  timeout¤	Ng  set-features¤	Og  droppable-feature¤	Pg  
set-lambda¤C 5h ,  ö   ] 4	 >  "  G   hÈ  f  ] 4>   "  G  4>   "  G  4>   "  G  4>   "  G  4>  "  G  4>  "  G   4	
>  "  G  4>   "  G  4>   "  G  4>   "  G  4>   "  G  4>  "  G  4>   "  G    4>   "  G  4>   "  G  4>  "  G  4>  "  G  4>   "  G   4>   "  G  4>   "  G  4>  "  G  4>  "  G  4>  "  G  4>   "  G    4>   "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>   "  G   4>   "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>   "  G    4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>   "  G   4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4
>  "  G  4>   "  G  		 C   ^      g  filenamef  
treize.scm
	
							#			3			C			U			X			]			i			j			l			m			p			w		 		 		  	 	 °	!	 À	"	 Ã	"	 È	"	 Ñ	$	 ä	%	 å	%	 ç	%	 ë	&	 ì	&	 î	&	 ï	'	 ÿ	(		)		)		)	 	*	#	*	(	*	1	,	D	-	E	-	G	-	H	.	X	/	h	0	k	0	p	0	y	1	|	1		1		2		2		2		4	®	5	¯	5	±	5	µ	6	¶	6	¸	6	¹	7	É	8	Ì	8	Ñ	8	Ú	9	Ý	9	â	9	ë	:	î	:	ó	:	ü	;	ÿ	;		;		=	 	>	!	>	#	>	$	?	4	@	7	@	<	@	E	A	H	A	M	A	V	B	Y	B	^	B	g	C	j	C	o	C	x	D	{	D		D		F		G		G		G	£	H	¤	H	¦	H	§	I	ª	I	¯	I	¸	J	»	J	À	J	É	K	Ì	K	Ñ	K	Ú	L	Ý	L	â	L	ë	M	î	M	ó	M	ü	N	ÿ	N		N		P	 	Q	!	Q	#	Q	$	R	'	R	,	R	5	S	8	S	=	S	F	T	I	T	N	T	W	U	Z	U	_	U	h	V	k	V	p	V	y	W	|	W		W		X		X		X		Z	 	Z	¥	Z	®	]	Ä	_	 	Å
  g  nameg  new-game CR        h   l   ] 45 6     d       g  filenamef  
treize.scm
	a
		b			b	 		
  g  nameg  give-status-message CR!"#$%&'       h    ¬   ] 45444
5556 ¤       g  filenamef  
treize.scm
	d
		e				e			e			e	"		f			f	!		f	)		f	!		f			e	 		
  g  nameg  get-stock-no-string C R()*+&   h8   Ü   ]4 5$  C45$  C4 
5$  
45CCÔ       g  slot-id
		8 g  	card-list		8  g  filenamef  
treize.scm
	h
		i			i			j			j			j			j			i		!	k		,	i		-	l	
	5	l	 		8	  g  nameg  button-pressed C,R(   h  ·  ]
 $  "   	$  C 
$  C 	$  ,4	5$  4	5$  	$  C	CCC 	$  ,4	5$  4	5$  	$  C	CCC 	$  ,4	5$  4	5$  	$  C	CCC 	$  ,4	5$  4	5$  	$  C	CCC 	$  ,4	5$  4	5$  	$  C	CCC 	$  ,4	5$  4	5$  	$  C	CCC 	$  ,4	5$  4	5$  	$  C	CCC 	$  ,4	5$  4	5$  	$  C	CCC 	$  ,4	5$  4	5$  	$  C	CCC 	$  ,4	5$  4	5$  	$  C	CCC 	$  ,4	5$  4	5$  	$  C	CCC 	$  ,4	5$  4	5$  	$  C	CCC 	$  ,4	5$  4	5$  	$  C	CCC 	
$  ,4	5$  4	5$  	$  C	CCC 		$  ,4	5$  4	5$  	$  C	CCC 	$  ,4	5$  4	5$  	$  C	CCC 	$  ,4	
5$  4	5$  	
$  C	CCC 	$  ,4		5$  4	
5$  		$  C	
CCC 	$  ,4	5$  4	5$  	$  C	CCC 	$  ,4	5$  4	5$  	$  C	CCC 	$  ,4	5$  4	5$  	$  C	CCCC  ¯      g  slot-id
	 g  r-slot	 g  t			  g  filenamef  
treize.scm
	n
		o			o				p			o		%	r			)	o		0	t			4	o		5	u		?	u			@	v		J	u			O	w		S	u			Z	x		[	x		e	y			i	o		j	z		t	z			u	{			z		 	|	 	z		 	}	 	}	 	~		 	o	 		 ©			 ª 	 ´			 ¹ 	 ½			 Ä 	 Å 	 Ï 		 Ó	o	 Ô 	 Þ 		 ß 	 é 		 î 	 ò 		 ù 	 ú 	 			o		 	 		 	 		# 	' 		. 	/ 	9 		=	o	> 	H 		I 	S 		X 	\ 		c 	d 	n 		r	o	s 	} 		~ 	 		 	 		 	 	£ 		§	o	¨ 	² 		³ 	½ 		Â 	Æ 		Í 	Î 	Ø 		Ü	o	Ý 	ç 		è 	ò 		÷ 	û 		  	  	 ¡			o	 ¢	 ¢		 £	' ¢		, ¤	0 ¢		7 ¥	8 ¥	B ¦		F	o	G §	Q §		R ¨	\ §		a ©	e §		l ª	m ª	w «		{	o	| ¬	 ¬		 ­	 ¬		 ®	 ¬		¡ ¯	¢ ¯	¬ °		°	o	± ±	» ±		¼ ²	Æ ±		Ë ³	Ï ±		Ö ´	× ´	á µ		å	o	æ ¶	ð ¶		ñ ·	û ¶		  ¸	 ¶		 ¹	 ¹	 º			o	 »	% »		& ¼	0 »		5 ½	9 »		@ ¾	A ¾	K ¿		O	o	P À	Z À		[ Á	e À		j Â	n À		u Ã	v Ã	 Ä			o	 Å	 Å		 Æ	 Å		 Ç	£ Å		ª È	« È	µ É		¹	o	º Ê	Ä Ê		Å Ë	Ï Ê		Ô Ì	Ø Ê		ß Í	à Í	ê Î		î	o	ï Ï	ù Ï		ú Ð	 Ï			 Ñ	 Ï		 Ò	 Ò	 Ó		#	o	$ Ô	. Ô		/ Õ	9 Ô		> Ö	B Ô		I ×	J ×	T Ø		X	o	Y Ù	c Ù		d Ú	n Ù		s Û	w Ù		~ Ü	 Ü	 Ú		  g  nameg  
available? C+R(+)-  h8     ]45$  C4 5$  	454455CC þ       g  
start-slot
		7 g  	card-list		7 g  end-slot			7  g  filenamef  
treize.scm
 Þ
	 ß		 ß		 à		 ß		 á		$ á		& á		' â		* â		2 â		3 á		4 á	 		7	  g  nameg  
droppable? C.R./0('12     h   Ì  ]4 5$  4	5$  w45$  j $  "  $  L4	5$  C4	5454 >  "  G  	44556CCCC Ä      g  
start-slot
	  g  	card-list	  g  end-slot		  g  t		+	? g  new-contents		W  g  moving-back		a	|  g  filenamef  
treize.scm
 ä
	 å		 å		 æ		 å		 ç		' å		+ è		+ è		< é		C è		D ê		N ê		Q ì	%	W ì		Z í	+	a í	&	a í		d î		l î	#	q î	  ï	!  ï	/  ï	*  ï	!  ï	 	 	  g  nameg  button-released C3R(45+)-*0'12/   hè   :  ] 
$  045$  4
5$  C
64
5$  C
64 5$  C4 
5$  44 55$  x4	 5$  k $  W4	5$  "  C4
	5454 >  "  G  4	44555"  $  6CCCC      2      g  slot-id
	 â g  new-contents  Ì g  moving-back	  µ  g  filenamef  
treize.scm
 ó
	 ô		
 ô		 õ	
	 õ		 ö		 ö	
	% ÷	%	' ÷		( ø		1 ø	
	8 ù	%	: ù		; ú		E ú		H û		S ú		T ü		W ü		_ ü		b ü		f ú		g þ		q ú		u ÿ		y ÿ		z 	  	 	) 	 	/ 	* 	 	 ¥	' ª	 ¶	 »	% ¾	3 Å	. Ç	% É	 Õ ú	 Ú	 -	 â  g  nameg  button-clicked C6R h   u   ]C    m       g  slot-id
		  g  filenamef  
treize.scm

 		  g  nameg  button-double-clicked C7R89      h(   ~   ] 4>   "  G  45 $  C6        v       g  filenamef  
treize.scm

							!	 		!
  g  nameg  game-continuable C:R( h       ] 4	5$  45$  
6CCw       g  filenamef  
treize.scm

										 		 
  g  nameg  game-won C8R;"<=>?@AB 
       h@   Ù   ] &  6 &  6 &  6 &  	6C    Ñ       g  suit
		<  g  filenamef  
treize.scm

	
												&		*		,		4		8		:	 		<  g  nameg  hint-remove-king CCR(+D*)-ECFG    hð   b  ]4 5$  "  	4 
5$   	$   	 6C44 55$   44	4 555645$  "  64
5$  "  	44 554455$  )	$  	 6 	$   	 6C
 6      Z      g  slot1
	 ê g  slot2	 ê g  t				" g  t		n ´ g  t	  ±  g  filenamef  
treize.scm

		
							
	&		+	
	/		4		9	"	;	
	@!		C!		K!		L!	
	P!		U"		X"	.	["	8	c"	.	e"		g"	
	h#		n#		|$	 $	 #	 %	  %	+ %	   &	  £&	+ «&	  ¬%	 ­%	 ®%	 ¸#	
 ½'	 Á'	 È(	$ Ê(	 Ï)	 Ó)	 Ø*	" Ý*	. ß*	 ê,	 0	 ê	  g  nameg  
check-move CDR("H h       ] 4
5$  C
45 C             g  filenamef  
treize.scm
.
	/		/		0		0		0		0	 		
  g  nameg  	dealable? CIR(&')-G    hP   ÷   ] 4	5$  C44	55$  -	44	5544	55$  			6CCï       g  filenamef  
treize.scm
3
	4		4		5	
	5		5	
	5		!4		$6		'6		/6		07		37	!	:7		=7		>6		?6		C4		L8	 		P
  g  nameg  check-waste CJRDJI    h0      ]4	5  $   C45   $   C6               g  t
	
	) g  t
		)  g  filenamef  
treize.scm
:
	;		
;		<		;		)=	 		)
  g  nameg  get-hint C9R      h   U   ] C    M       g  filenamef  
treize.scm
?
 		
  g  nameg  get-options CKR      h   m   ]C    e       g  options
		  g  filenamef  
treize.scm
B
 		  g  nameg  apply-options CLR      h   Q   ] C    I       g  filenamef  
treize.scm
E
 		
  g  nameg  timeout CMR4NiOi>  "  G  Pii,i3i6i7i:i8i9iKiLiMi.i6î       g  filenamef  
treize.scm		
	y	

	a

ÿ	d
/	h
	n
Ù Þ
_ ä
 ­ ó
!8
!ù
"ª
#ì
'e
($.
)3
*p:
*à?
+hB
+ÔE
+ÕH
, J
 	, 
   C6 