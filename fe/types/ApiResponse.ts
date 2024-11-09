// types/ApiResponse.ts

import { Transaction } from "./Transaction";

export interface ApiResponse {
  total: number;
  transactions: Transaction[];
  page: number;
  total_elements: number;
}
