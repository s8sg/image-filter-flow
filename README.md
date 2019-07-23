# image-filter-flow
Flow that apply filter effect on resized image

## Dependency
### Filters
##### coherent-line-drawing
> source: https://github.com/esimov/openfaas-coherent-line-drawing       
Need input to be in image mode
    
    
> Use more filter just by changing the filter query

### Image Resizer - ResizeImageMagick
> Source: https://github.com/openfaas/faas/tree/master/sample-functions/ResizeImageMagick   


## Getting started 
Clone
```
https://github.com/s8sg/image-filter-flow.git
```
Build 
```
faas template pull https://github.com/alexellis/opencv-openfaas-template
faas template pull https://github.com/s8sg/faas-flow 
faas build
```
Deploy
```
faas deploy
```
Invoke
```
cat openfaas-logo.jpg | faas invoke image-filter-flow --query="filter=coherent-line-drawing" > output.jpg
```
