export const paymentData = {
  object: "event",
  id: "evnt_test_5vgegbyf02uvuz1gii7",
  livemode: false,
  location: "/events/evnt_test_5vgegbyf02uvuz1gii7",
  webhook_deliveries: [],
  data: {
    object: "charge",
    id: "chrg_test_5vgfbi0yj1nc9rxzxp5",
    location: "/charges/chrg_test_5vgegbvfx7qrd5peb9k",
    amount: 100000,
    authorization_type: "",
    authorized_amount: 100000,
    captured_amount: 100000,
    net: 98234,
    fee: 1650,
    fee_vat: 116,
    interest: 0,
    interest_vat: 0,
    funding_amount: 100000,
    refunded_amount: 0,
    transaction_fees: {
      fee_flat: "0.0",
      fee_rate: "1.65",
      vat_rate: "7.0",
    },
    platform_fee: {
      fixed: "",
      amount: "",
      percentage: "",
    },
    currency: "THB",
    funding_currency: "THB",
    ip: "",
    refunds: {
      object: "list",
      data: [],
      limit: 20,
      offset: 0,
      total: 0,
      location: "/charges/chrg_test_5vgegbvfx7qrd5peb9k/refunds",
      order: "chronological",
      from: "1970-01-01T00:00:00Z",
      to: "2023-04-15T11:32:08Z",
    },
    link: "",
    description: "",
    metadata: {},
    card: "",
    source: {
      object: "source",
      id: "src_test_5vgegbtpa9ffpijmfiy",
      livemode: false,
      location: "/sources/src_test_5vgegbtpa9ffpijmfiy",
      amount: 100000,
      barcode: "",
      bank: "",
      created_at: "2023-04-15T11:32:08Z",
      currency: "THB",
      email: "",
      flow: "offline",
      installment_term: "",
      absorption_type: "",
      name: "",
      mobile_number: "",
      phone_number: "",
      platform_type: "",
      scannable_code: {
        object: "barcode",
        type: "qr",
        image: {
          object: "document",
          livemode: false,
          id: "docu_test_5vgegbxaghghibi0v7t",
          deleted: false,
          filename: "qrcode_test.svg",
          location:
            "/charges/chrg_test_5vgegbvfx7qrd5peb9k/documents/docu_test_5vgegbxaghghibi0v7t",
          kind: "qr",
          download_uri:
            "https://api.omise.co/charges/chrg_test_5vgegbvfx7qrd5peb9k/documents/docu_test_5vgegbxaghghibi0v7t/downloads/CF75F55B445D3A7E",
          created_at: "2023-04-15T11:32:08Z",
        },
      },
      billing: "",
      shipping: "",
      items: [],
      references: "",
      store_id: "",
      store_name: "",
      terminal_id: "",
      type: "promptpay",
      zero_interest_installments: "",
      charge_status: "pending",
      receipt_amount: "",
      discounts: [],
    },
    schedule: "",
    customer: "",
    dispute: "",
    transaction: "",
    failure_code: "",
    failure_message: "",
    status: "pending",
    authorize_uri: "",
    return_uri: "",
    created_at: "2023-04-15T11:32:08Z",
    paid_at: "",
    expires_at: "2023-04-16T11:32:08Z",
    expired_at: "",
    reversed_at: "",
    zero_interest_installments: false,
    branch: "",
    terminal: "",
    device: "",
    authorized: false,
    capturable: false,
    capture: true,
    disputable: false,
    livemode: false,
    refundable: false,
    reversed: false,
    reversible: false,
    voided: false,
    paid: false,
    expired: false,
  },
  key: "charge.create",
  created_at: "2023-04-15T11:32:08Z",
};
