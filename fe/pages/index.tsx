import { useState, useEffect, FormEvent } from "react";
import axios from "axios";
import SearchResults from "@/components/SearchResults";
import { ApiResponse } from "@/types/ApiResponse";
import styles from "@/styles/Home.module.css";

const Home = () => {
  const [searchQuery, setSearchQuery] = useState<string>("");
  const [transactions, setTransactions] = useState<ApiResponse["transactions"]>(
    []
  );
  const [page, setPage] = useState<number>(1);
  const [totalPages, setTotalPages] = useState<number>(1);
  const [loading, setLoading] = useState<boolean>(false);

  useEffect(() => {
    if (searchQuery) {
      const fetchData = async () => {
        await fetchTransactions(searchQuery, page);
      };
      fetchData();
    }
  }, [searchQuery, page]);

  const fetchTransactions = async (query: string, page: number) => {
    setLoading(true);
    try {
      let params: any = { page: page, page_size: 10, search: query };

      const response = await axios.get<ApiResponse>(
        "http://localhost:8080/get-transactions",
        { params }
      );
      setTransactions(response.data.transactions);
      setTotalPages(Math.ceil(response.data.total / 10)); // Assuming your API returns the total records
    } catch (error) {
      if (axios.isAxiosError(error)) {
        console.error("Network error:", error.message);
      } else {
        console.error("Error fetching transactions:", error);
      }
    }
    setLoading(false);
  };

  const handleSearch = (e: FormEvent) => {
    e.preventDefault();
    setPage(1); // Reset page number on new search
    fetchTransactions(searchQuery, 1);
  };

  return (
    <div className={styles.container}>
      <h1 className={styles.title}>Search Transactions</h1>
      <form onSubmit={handleSearch} className={styles.searchForm}>
        <input
          type="text"
          placeholder="Search by date, transaction no, credit, debit, or detail"
          value={searchQuery}
          onChange={(e) => setSearchQuery(e.target.value)}
          className={styles.input}
        />
        <button type="submit" className={styles.button}>
          Search
        </button>
      </form>

      {loading ? (
        <p className={styles.loading}>Loading...</p>
      ) : (
        <>
          <SearchResults transactions={transactions} />
          <div className={styles.pagination}>
            <button
              onClick={() => setPage(page > 1 ? page - 1 : 1)}
              disabled={page === 1}
              className={styles.paginationButton}
            >
              Previous
            </button>
            <span className={styles.pageInfo}>
              {" "}
              Page {page} of {totalPages}{" "}
            </span>
            <button
              onClick={() => setPage(page < totalPages ? page + 1 : totalPages)}
              disabled={page === totalPages}
              className={styles.paginationButton}
            >
              Next
            </button>
          </div>
        </>
      )}
    </div>
  );
};

export default Home;
