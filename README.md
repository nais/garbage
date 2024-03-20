# garbage

*Google Artifact Registry Brutally Automatic Garbage Eradicator* (Garbage) deletes unused images from Google Artifact Registry (GAR).

Garbage periodically cross-references images in GAR with resources in Kubernetes and notes which images are in use.
After a certain grace period, images no longer in use are deleted.

## Detailed algorithm

- Query the Google Artifact Registry (GAR) API to fetch all image metadata in a given project
- Store the image metadata (id, name, tag, shasum) in a database
- Query all Kubernetes clusters through [nais/api](https://github.com/nais/api) to fetch all image references
  - From `ReplicaSet`, `Job`, `CronJob`, `StatefulSet`, `DaemonSet`.
- Generate deletion candidates using the garbage collection approach (collect and sweep):
  - Initialize empty hashmap with all images from GAR. Set `refcount` to 0.
  - Cross-reference images from Kubernetes with GAR images. Increment `refcount` by 1 for each match.
  - For each image with `refcount>0`, recursively find parent references and increment `refcount` by 1 for each match.
  - All images with `refcount>0` are marked as seen in the database with a timestamp.
- Delete images that haven't been seen recently (within a given threshold, e.g. 30 days)
