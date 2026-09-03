"""INTENTIONALLY FLAWED Spark transform for candidates to diagnose."""
from pyspark.sql import SparkSession

spark = SparkSession.builder.appName("evidence").getOrCreate()
rows = spark.read.json("/data/evidence/*.json")

# Faults: collect moves an unbounded dataset to the driver; no schema,
# checkpoint, event time, deduplication, quarantine, or deterministic output.
all_rows = rows.collect()
spark.createDataFrame(all_rows).coalesce(1).write.mode("overwrite").json("/data/latest")

