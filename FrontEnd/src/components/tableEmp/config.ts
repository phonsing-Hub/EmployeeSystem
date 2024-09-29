const columns = [
  { name: "ID", uid: "id", sortable: true },
  { name: "NAME", uid: "firstname", sortable: true },
  { name: "ROLE", uid: "role" },
  { name: "CONTACT", uid: "contact" },
  { name: "SALARY", uid: "salary", sortable: true },
  { name: "ACTIONS", uid: "actions" },
];

const formatCurrencyTHB = (amount: number): string => {
  return amount.toLocaleString("th-TH");
};

const formatDate = (input: string): string => {
  if (input === "") return "";
  const date = new Date(input);
  return date.toISOString().split("T")[0];
};

const formatPhoneNumber = (phoneNumber: string): string => {
  return phoneNumber.replace(/\./g, '-');
}
export { columns, formatCurrencyTHB, formatDate, formatPhoneNumber };
