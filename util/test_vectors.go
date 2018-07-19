package util

type TestVector struct {
	Magic int
    Hrp string
    EncodedTxref string
    Height int
    Position int
    EncFail int //0 == must not fail, 1 == can fail, 2 == can fail and continue with next test, 3 == skip
    DecFail int
    NonStd bool
}

var TestVectors = []TestVector{
	{
        TxrefMagicBtcMainnet,
        TxrefBech32HrpMainnet,
        "tx1:rqqq-qqqq-qmhu-qk",
        0,
        0,
        0,0,false,
    },
	
	{
        TxrefMagicBtcMainnet,
        TxrefBech32HrpMainnet,
        "tx1:rjk0-u5ng-4jsf-mc",
        466793,
        2205,
        0,0,false,
    },
	
    {
        TxrefMagicBtcMainnet,
        TxrefBech32HrpMainnet,
        "tx1:rjk0-u5n1-2jsi-mc", /* error correct test >2tsi< instead of >4jsf<*/
        466793,
        2205,
        1,0,false,
    },
	
    {
        TxrefMagicBtcMainnet,
        TxrefBech32HrpMainnet,
        "tx1:rqqq-qqqq-qmhu-qk",
        0,
        0,
        0,0,false,
    },
	
    {
        TxrefMagicBtcMainnet,
        TxrefBech32HrpMainnet,
        "tx1:rzqq-qqqq-uvlj-ez",
        1,
        0,
        0,0,false,
    },
	
    {
        TxrefMagicBtcMainnet,
        TxrefBech32HrpMainnet,
        "tx1:rjk0-u5ng-4jsf-mc", /* complete invalid */
        0,
        0,
        1,1,false, /* enc & dec must fail */
    },
	
    {
        TxrefMagicBtcMainnet,
        TxrefBech32HrpMainnet,
        "tx1:r7ll-lrar-a27h-kt",
        2097151, /* last valid block height with current enc/dec version is 0x1FFFFF*/
        1000,
        0,0,false,
    },
	
    {
        TxrefMagicBtcMainnet,
        TxrefBech32HrpMainnet,
        "", /* encoding must fail, no txref to chain against */
        2097152, /* invalid height */
        1000,
        2,1,false,
    },
	
    {
        TxrefMagicBtcMainnet,
        TxrefBech32HrpMainnet,
        "tx1:r7ll-llll-khym-tq",
        2097151, /* last valid block height with current enc/dec version is 0x1FFFFF*/
        8191, /* last valid tx pos is 0x1FFF */
        0,0,false,
    },
	
    {
        TxrefMagicBtcMainnet,
        TxrefBech32HrpMainnet,
        "", /* encoding must fail, no txref to chain against */
        2097151, /* last valid block height with current enc/dec version is 0x1FFFFF*/
        8192, /* invalid tx pos */
        2,1,false,
    },
	
    {
        TxrefMagicBtcMainnet,
        TxrefBech32HrpMainnet,
        "tx1:r7ll-lrqq-vq5e-gg",
        2097151, /* last valid block height with current enc/dec version is 0x1FFFFF*/
        0,
        0,0,false,
    },
	
    {
        TxrefMagicBtcMainnet,
        TxrefBech32HrpMainnet,
        "tx1:rqqq-qull-6v87-r7",
        0,
        8191, /* last valid tx pos is 0x1FFF */
        0,0,false,
    },
	
    {
        TxrefMagicBtcMainnet,
        TxrefBech32HrpMainnet,
        "tx1:rjk0-u5ng-gghq-fkg7", /* valid Bech32, but 10x5bit packages instead of 8 */
        0,
        0,
        3,2,false, /* ignore encoding */
    },
	
    {
        TxrefMagicBtcMainnet,
        TxrefBech32HrpMainnet,
        "tx1:rjk0-u5qd-s43z", /* valid Bech32, but 6x5bit packages instead of 8 */
        0,
        0,
        3,2,false, /* ignore encoding */
    },
	
    {
        0xB,
        TxrefBech32HrpMainnet,
        "tx1:t7ll-llll-gey7-ez",
        2097151,
        8191,
        0,0,false, /* ignore encoding */
    },
	
    {
        TxrefMagicBtcMainnet,
        TxrefBech32HrpMainnet,
        "tx1:rk63-uvxf-9pqc-sy",
        467883,
        2355,
        0,0,false, /* ignore encoding */
    },
	
    {
        TxrefMagicBtcTestnet,
        TxrefBech32HrpTestnet,
        "txtest1:xk63-uqvx-fqx8-xqr8",
        467883,
        2355,
        0,0,true, /* ignore encoding */
    },
	
    {
        TxrefMagicBtcTestnet,
        TxrefBech32HrpTestnet,
        "txtest1:xqqq-qqqq-qqkn-3gh9",
        0,
        0,
        0,0,true, /* ignore encoding */
    },
	
    {
        TxrefMagicBtcTestnet,
        TxrefBech32HrpTestnet,
        "txtest1:x7ll-llll-llj9-t9dk",
        0x3FFFFFF,
        0x3FFFF,
        0,0,true, /* ignore encoding */
    },
}

