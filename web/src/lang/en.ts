export default {
  //通用
    com:{
      btn:{
        confirm:"Confirm",
        cancel:"Cancel",
        reset:"Reset",
        choose_file:"Choose File",
        connect :"Connect Wallet",
      },
      address:"Address",
      balance:"Balance",
      msg:{
        wrong_addr:"Please enter the correct address",
        wrong_hash:"Please enter the correct hash",
        wrong_pvk:"Please enter the correct private key",
      },
    },

    // 路由国际化
    route: {
      unavailable  :'Unavaliable',
      liveMint     : 'Live Mint',
      quickSell    : 'Quick Sell',

      nft          : 'NFT',
      freeMint     : 'Free Mint',
      followMint   : 'Follow Mint',
      hashMint     : 'Hash Mint',
      quickOffer   : 'Quick Offer',
      sweepFloor   : 'Sweep Floor',

      erc20        : 'ERC20',
      hotToken     : 'Hot Token',
      bigFlow      : 'Big Flow',
      tokenAnalysis: 'Token Analysis',
      walletAnalysis : 'Wallet Analysis',

      commonTools   : 'Common Tools',
      frontRun     : 'Front Run',
      tokenManage  : 'Token Manage',
      nftManage    : 'NFT Manage',

      userFeedBack: 'Feedback',
      
      superAdmin  :"Administrator",
      menu        :"Menu Manage",
      api         :"Api Manage",
      user        :"User Manage",
      authority   :"Role Manage",

      adminFeedBack: 'User Feedback',
      upgrade     :  "Upgrade",
      login       :"Login",
      404         :"Not Found",

      //此项为方便upgrade卡片直接写在这里
      newFeatures :"New features in the future"
    },

    // 代币分发 / 归集
    tm:{
      label:{
        mm_transfer:"Use Metamask:",
        etra_all:"Extract All:",
        main_wal:"Main Wallet:",
        main_wal_c:"Receiver Wallet:",
        main_wal_pvk:"Other Wallet?:",
        token_addr:"Token Address:",
        pri_fee :"Priority Fee:",
        max_fee:"Max Fee:",
        ava_bal:"Avaliable:",
        each_wal:"Each address to send:",
        transfer:"transfer in",
        transferout:"transfer out",
      },
      title:{
        dis:"Token MultiSender",
        coll:"Token Collection"
      },
      ph:{
        main_wal:"Please select a wallet as sender",
        main_wal_c:"Please select a wallet as receiver",
        pvk:"Please paste private key of sender wallet",
        pvk_c:"Please paste address of receiver wallet",
        token:"Please paste the token address (default is eth)",
      },
      btn:{
        choose_file:"Receive Address File",
        choose_file_c:"Sender Address File",

      },
      tip:{
        file:"One per line"
      }
    },

    setting:{
    title:"Global Configuration",
    mode:"Dark Mode",
    warn:"Note: The data will be reset after refreshing the page!",
    label:{
      node:"Customize Node",
      pvk:"Private Keys",
    },
    btn:{
      set:"Set",
      reset:"Restore",
      choose_file:"Choose File",
      clear_list:"Clear List",
      parse:"Parse",
    },
    ph:{
      http:"Please enter the node url with the api key",
      wss:"Please enter the node url with the api key",
      txt_pvk:"Paste or load private keys from a file, one per line or separated by commas",
      selected:"Selected :"
    },
    table:{

    }
  },
  upgrade:{
    title : "Choose Package",
    month : "/ Month",
    year  : "/ Year",
    tip   : "You can use the following features with this package.",
    buy   : "Upgrade",
    wait  : "Waiting...",
    cnf   : "Pending...",
    btn   : "Upgrade",
  },
  hashmint:{
    message:"It's recommended to operate with alt account when batching mint",
    title:{
      pvk:"Private Keys"
    },
    ph:{
      txtarea_pvk:"Paste or load private keys from a file, one per line or separated by commas",
      txtarea_data:"Automatic parsing",
      input_hash:"Paste the transaction hash, manual input is not recommended",
      input_address:"Automatically parse the contract address",
      input_value:"(ETH)",
      input_maxfee:"Max Gas Fee(GWei)",
      input_prifee:"Max priority (GWei)",
      input_gaslimit:"Gas limit",
      input_cost:"Max cost",
    },
    btn:{
      choose_file:"Select File",
      clear_list:"Clear List",
      parse_key:"Parse",
      reset:"Reset Form",
      mm_mint:"Metamask Mint",
      batch_mint:"Batch Auto mint",
      delete:"Delete",
    },
    ld:{
      table:"Parsing..."
    },
    label:{
      contract_addr:"Contract Addr",
      max_cost:"Max Cost",
      address:"Address",
      balance:"Balance",
      status:"Status",
      hash:"Tx Hash",
      op:"Operation",
      res:"Result",
    },
    tip:{
      hash:"Note: Find the hash value of someone else's successful transaction in EtherScan and copy it here.",
    },
    description:"Private key not imported",
    msg:{
      copy:"Copy Success!",
      hash:"Please enter the correct transaction hash!",
      up_type:"Only supports uploading .txt files!",
      up_size:"Upload file size cannot exceed 5KB!",
      up_fail:"Failed, please upload again",
      bal:" No balance wallet has been skipped (unable to participate in mint)!",
      w_pvk:"Your input contains an incorrect private key!",
      parse_f:"Parse finished!",
      not_send:"Not sent: Insufficient balance",
    },
  },

  //快速反馈-用户
  ufeedback:{
    ph:{
      type:"Please select",
      topic:"Please enter subject of your feedback",
      phone:"Please enter phone number (optional)",
      txt:"Please enter the suggestions or comments you want to give feedback... or you can send emails directly to: talktogang{'@'}gmail.com",
    },
    label:{
      type:"Feedback Type",
      topic:"Feedback Subject",
      phone:"Phone Number",
    },
    btn:{
      submit:"Submit"
    },
  },
  // 导航栏国际化
  navbar: {
    tip:{
      sev:"Contact Customer Service",
      full:"Full screen",
      exit:"Exit full screen",
      ref:"Refresh",
    },
   ph:{
      search:"Address Analysis <--Enter a address",
   },
  }
};
