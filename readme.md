# Overview

- This project provides a REST API for streaming and processing large CSV files.
- Inspired by the <a href="https://github.com/olist/work-at-olist">Olist-Challenge</a> and it's using a <a href="https://www.kaggle.com/datasets/rounakbanik/the-movies-dataset">Dataset</a> of movies to test upload. The API is designed to handle large datasets efficiently, allowing users to upload csv data into MongoDb.

---

## 🔥 **Key Features**

- **Large CSV File Uploads:** Efficient processing using streaming to avoid memory overload, with **Goroutines** for concurrent processing of data rows.
- **Data Validation and Parsing:** Verifies and transforms data before storing it in the database.
- **MongoDB Support:** Saves data directly into MongoDB collections.
- **RESTful Endpoints:** Simple interface to interact with the system.
- **Movie Dataset:** Example dataset used for testing uploads.

---

## 📊 **Test Dataset**
- The test dataset contains information about movies, such as title, genre, release year, and director. You can customize it or use your own CSV file.

```csv  
movieId, title,genres
1,Inception, Action| Sci-Fi,
2,Titanic,Drama| Romance
3,The Matrix,Action| Sci-Fi
```
