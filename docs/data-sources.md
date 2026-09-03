# Public data-source ideas

Verify the source's current terms and document your retrieval date. These links
lead to the authoritative publishers rather than third-party mirrors.

| Source | Shape | Example questions |
|---|---|---|
| [openFDA APIs](https://open.fda.gov/apis/) | Live JSON APIs and bulk downloads for drugs, devices, food, cosmetics, and recalls | Which recall categories or adverse-event signals are changing? |
| [EPA CompTox data and APIs](https://www.epa.gov/comptox-tools/computational-toxicology-and-exposure-apis) | Chemical, exposure, bioactivity, and toxicity data; API key required for some endpoints | Which evidence gaps should a safety team prioritize? |
| [NIEHS ICE](https://ice.ntp.niehs.nih.gov/) | Curated in vivo, in vitro, and in silico toxicology datasets | How do alternative assays compare for an endpoint? |
| [CDC Environmental Public Health Tracking](https://ephtracking.cdc.gov/) | Environmental hazards, exposures, and health outcomes by place/time | Where are environmental and health trends moving together? |
| [NCBI PubChem PUG REST](https://pubchem.ncbi.nlm.nih.gov/docs/pug-rest) | Chemical identifiers, properties, structures, and bioassay links | Can source identifiers be resolved and standardized reliably? |
| [NCBI E-utilities](https://www.ncbi.nlm.nih.gov/books/NBK25501/) | PubMed and other Entrez databases | What new evidence appeared for a substance or mechanism? |
| [USDA FoodData Central API](https://fdc.nal.usda.gov/api-guide.html) | Foods, nutrients, branded-food updates | How do formulations or nutrient profiles change over time? |
| [Data.gov](https://data.gov/) | Catalog spanning federal agencies | Find a high-value source we did not anticipate |

## Good project directions

- A live regulatory-change radar that normalizes recalls from one or more
  sectors and shows explainable trend changes.
- A BioBrick joining a difficult government dataset to stable chemical
  identifiers, with mapping-quality and drift reports.
- An evidence map connecting substances, assays, genes/pathways, and literature,
  with every edge traceable to a source record.
- A streaming early-warning dashboard with explicit event time, deduplication,
  backfill, and freshness behavior.
- A defensible baseline model for an endpoint, packaged behind an API with its
  applicability domain and uncertainty visible.

