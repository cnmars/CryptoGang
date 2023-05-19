export default {
    //通用
    com:{
      btn:{
        confirm:"确定",
        cancel:"取消",
        reset:"重置",
        choose_file:"选择文件",
        connect :"钱包登录",
      },
      address:"地址",
      balance:"余额",
      msg:{
        wrong_addr:"请输入正确的地址",
        wrong_hash:"请输入正确的Hash",
        wrong_pvk:"请输入正确的私钥",
      },
    },

  // 路由国际化
  route: {
      unavailable   :'未开放',
      liveMint      : 'Live Mint',
      quickSell     : '一键挂单',

      nft           : 'NFT',
      freeMint      : 'Free Mint',
      followMint    : 'Follow Mint',
      hashMint      : 'Hash Mint',
      quickOffer    : '一键Offer',
      sweepFloor    : '扫地板',

      erc20         : 'ERC20',
      hotToken      : '热门Token',
      bigFlow       : '大额监控',
      tokenAnalysis : '代币分析',
      walletAnalysis: '钱包分析',

      commonTools    : '常用工具',
      frontRun      : '抢跑机器人',
      tokenManage   : '资金管理',
      nftManage     : 'NFT管理',

      userFeedBack  : '快速反馈',

      superAdmin    :"管理员",
      menu          :"菜单管理",
      api           :"Api管理",
      user          :"用户管理",
      authority     :"角色管理",
      login         :"签名登录",
      adminFeedBack : '用户反馈',
      upgrade       : "升级账户",
      404           :"Not Found",
      //此项为方便upgrade卡片直接写在这里
      newFeatures   :"后续上架的新功能",
  },

    // 代币分发 / 归集
    tm:{
      label:{
        mm_transfer:"小狐狸发送:",
        etra_all:"全部提取:",
        main_wal:"发送钱包:",
        main_wal_c:"接收钱包:",
        main_wal_pvk:"别的钱包？:",
        token_addr:"代币地址:",
        pri_fee :"小费:",
        max_fee:"Max Fee:",
        ava_bal:"可用余额:",
        each_wal:"每个钱包发送:",
        transfer:"转入",
        transferout:"转出",
      },
      title:{
        dis:"代币分发",
        coll:"代币归集"
      },
      ph:{
        main_wal:"请选择一个钱包作为发送方",
        main_wal_c:"请选择一个钱包作为接收方",
        pvk:"假如使用别的钱包则粘贴其私钥",
        pvk_c:"假如使用别的钱包接收请粘贴地址",
        token:"请粘贴代币合约地址(默认为ETH)",
      },
      btn:{
        choose_file:"接收地址文件",
        choose_file_c:"子钱包私钥文件",
      },
      tip:{
        file:"每行一个"
      }
    },

  //右侧设置栏
  setting:{
    title:"全局配置",
    mode:"暗黑模式",
    warn:"注意: 刷新网页后数据将重置！",
    label:{
      node:"自定义节点",
      pvk:"私钥",
    },
    ph:{
      http:"请输入带有Api Key的节点连接(或自建节点连接)",
      wss:"请输入带有Api Key的节点连接(或自建节点连接)",
      txt_pvk:"粘贴或者从文件中加载私钥，一行一个或者以逗号隔开",
      selected:"已选取 :"
    },
    btn:{
      set:"设置",
      reset:"恢复默认",
      choose_file:"选取文件",
      clear_list:"清空列表",
      parse:"解析",
    },
  },
  upgrade:{
    title : "选择套餐",
    month : "/ 每月",
    year  : "/ 每年",
    tip   : "购买此套餐您将享有以下功能",
    buy   : "购买",
    wait  : "等待支付中...",
    cnf   : "交易确认中...",
    btn   : "账户升级",
  },
  //hash mint
  hashmint:{
    message:"Hash Mint中批量mint时建议用小号操作",
    title:{
      pvk:"私钥"
    },
    ph:{
      txtarea_pvk:"粘贴或者从文件中加载私钥，一行一个或者以逗号隔开",
      txtarea_data:"自动解析",
      input_hash:"粘贴交易哈希,不建议手动输入",
      input_address:"自动解析合约地址",
      input_value:"(ETH)",
      input_maxfee:"最大Gas Fee(GWei)",
      input_prifee:"最大小费(GWei)",
      input_gaslimit:"gas限制",
      input_cost:"最大成本",
    },
    btn:{
      choose_file:"选取文件",
      clear_list:"清空列表",
      parse_key:"解析",
      reset:"重置数据",
      mm_mint:"小狐狸跟单",
      batch_mint:"多号跟单",
      delete:"删除",
    },
    ld:{
      table:"解析中..."
    },
    tip:{
      hash:"注：在EtherScan区块浏览器中找到别人交易成功的哈希值复制到此处。",
    },
    label:{
      contract_addr:"合约地址",
      max_cost:"最大成本",
      address:"地址",
      balance:"余额",
      status:"状态",
      hash:"交易Hash",
      op:"操作",
      res:"结果",
    },
    description:"未导入私钥",
    msg:{
      copy:"复制成功！",
      hash:"请录入正确的交易hash！",
      up_type:"仅支持上传.txt文件",
      up_size:"上传文件大小不能超过 5KB!",
      up_fail:"失败，请重新上传",
      bal:"已略过没有余额的钱包(无法参与mint)！",
      w_pvk:"您的录入含有不正确的私钥！",
      parse_f:"解析完成！",
      not_send:"未打出：余额不足",
    },
  },
  ufeedback:{
    ph:{
      type:"请选择",
      topic:"请输入反馈的主题",
      phone:"请输入联系电话(可选)",
      txt:`请录入您想要反馈的建议或意见....或者您可以直接发送邮件到:talktogang{'@'}gmail.com`,
    },
    label:{
      type:"反馈类别",
      topic:"反馈主题",
      phone:"联系电话",
    },
    btn:{
      submit:"提交"
    },
  },

  navbar: {
    tip:{
      sev:"联系客服",
      full:"全屏",
      ref:"刷新",
    },
   ph:{
      search:"地址分析 <--输入地址",
   },
  }
};
