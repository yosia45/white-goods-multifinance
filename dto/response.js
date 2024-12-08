const getAllPurchaseData = [
  {
    purchase_id: "",
    monthly_payment: "",
    is_completed: false,
    item_tenors: [
      {
        item_tenord_id: "",
        interest: "",
        item: {
          item_id: "",
          name: "",
          on_the_road: {
            otr_id: "",
            name: "",
          },
        },
      },
    ],
    tenor: {
      tenor_id: "",
      duration: "",
    },
  },
];

const purchaseByID = {
  purchase_id: "",
  monthly_payment: "",
  is_completed: false,
  item_tenors: [
    {
      item_tenord_id: "",
      interest: "",
      item: {
        item_id: "",
        name: "",
        on_the_road: {
          otr_id: "",
          name: "",
        },
      },
    },
  ],
  tenor: {
    tenor_id: "",
    duration: "",
  },
  transactions: [
    {
      transaction_id: "",
      total_amount: "",
      payment_date: "",
      invoice_number: "",
      status: "",
    },
  ],
  customer: {
    customer_id: "",
    full_name: "",
    legal_name: "",
    limit_information: {
        user_limit_id: "",
        limit: "",
        current_limit: "",
    }
  }
};
