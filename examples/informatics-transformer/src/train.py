"""INTENTIONALLY FLAWED Tox21 modeling sketch — not a reference solution."""
import pandas as pd
from sklearn.cluster import KMeans
from sklearn.model_selection import train_test_split
from sklearn.preprocessing import StandardScaler

frame = pd.read_parquet("tox21.parquet")

# Faults: missing assay labels become negatives; scaler sees all data; random
# row split leaks duplicate/scaffold relatives; no baseline or held-out test.
y = frame.filter(regex="^assay_").fillna(0)
x = StandardScaler().fit_transform(frame.filter(regex="^feature_"))
x_train, x_test, y_train, y_test = train_test_split(x, y, random_state=1)

# Fault: cluster count is asserted without stability or chemical validation.
clusters = KMeans(n_clusters=8, random_state=1).fit_predict(x)
frame.assign(cluster=clusters).to_csv("clusters.csv", index=False)

