# Serif Health Takehome Interview

This repository contains the files and instructions for our takehome engineering interview. Please *locally* copy to your own public repo or [import](https://github.com/new/import) to your github account for use in sharing solutions back to us. Direct public forks and pull requests will expose your identity and solution to other candidates also working on this interview question, and we want the interview process to be fair for everyone. 

## Context
Serif Health was founded with a mission to make the US healthcare system more transparent, efficient, and affordable for everyone. One of the challenging problems we're solving at Serif Health is making healthcare *pricing* data transparent and uniform for all market participants. There are myriad reasons this is difficult.

At the macro level:
- Most data in healthcare is protected by law, sensitive by default and tends to be locked up in proprietary systems or data formats.
- The data aggregators and clearinghouses that do have access to clean, normalized bulk data assets tend to employ extremely expensive and restrictive licensing terms. 
- While recent price transparency laws have required hospitals and carriers to publish their pricing, compliance and data sharing occurs at varying levels of completeness and consistency.  

At the micro level:
- Medical billing and coding for a specific procedure can be very complicated and is contingent on place of service, patient history and comorbidities, structure of insurance arrangements, so on and so forth. Many procedures are lots of N of 1 type cases. 
- Insurance companies (carriers) establish pre-negotiated non-published contracted rates with each facility, physician group, or health system that reimburses the healthcare provider at a rate and structure very different from what is 'charged'. 

Summed together, all this complexity contributes to a general lack of transparency and market efficency in our healthcare system.



## Objective
In July 2022, insurance carriers were required to publish their negotiated prices with all providers and facilities under the Transparency in Coverage Act. Pricing for every procedure code for every provider in the country is a lot of data; thus, the published files are extremely large and require some forethought and skill to be able to work with them. 

Our customers typically want to know and compare reimbursement rates for healthcare services from specific carriers. E.g., what does Anthem reimburse orthopedic surgeons in New York state for total knee replacement surgery? To get there, we need to go to Anthem's Transparency in Coverage website, find their appropriate index file (also called a table of contents file), look up the MRF file URLs in the index for the correct plan, pull the MRF, extract the data, and we have our answer. The challenge for us is that carriers don't always follow the schemas, so these postings and indexes aren't always easy to decipher - it takes some sleuthing and creativity to get to the answers we seek. 

For this interview, we'll give you an index file URL and we'll skip in-network MRF processing for now, since the data elements in the in-network file are significantly more complex and variant. 

Your task is to write some code that can open an index file, stream through it, and isolate a set of network files in the index. We'd simply like to know, *what is the list of machine readable file URLs that represent the Anthem PPO network in New York state*? 


## Inputs
The input to this takehome is the Anthem machine readable index file [table of contents](https://antm-pt-prod-dataz-nogbd-nophi-us-east1.s3.amazonaws.com/anthem/2024-04-01_anthem_index.json.gz) for April 2024. 

You should write code that can open the machine readable index file and extract some in-network file URLs from it according to the schema published at [CMS' transparency in coverage repository](https://github.com/CMSgov/price-transparency-guide/tree/master/schemas/table-of-contents), so you can extract the data requested.

## Outputs
Your output should be the list of machine readable file URLs corresponding to Anthem's PPO in New York state. Make sure to read through the hints and pointers section before declaring your solution complete.

## Hints and Pointers
As you start working with the index, you'll quickly notice that the index file itself is extremly large, data is very frequently repeated, plan descriptions seem to contain random businesses in various regions around the country, and that there are a handful of different url styles. 

- How do you handle the file size and format efficiently, when the uncompressed file will exceed memory limitations on most systems? 
- When you look at your output URL list, which segments of the URL are changing, which segments are repeating, and what might that mean?
- Is the 'description' field helpful? Is it complete? Does it change relative to 'location'? Is Highmark the same as Anthem?
- Anthem has an interactive MRF lookup system. This lookup can be used to gather additional information - but it requires you to input the EIN or name of an employer who offers an Anthem health plan: [Anthem EIN lookup](https://www.anthem.com/machine-readable-file/search/). How might you find a business likely to be in the Anthem NY PPO? How can you use this tool to confirm if your answer is complete?

Use creative thinking and your best judgement to proceed here, and discuss your decisions in your writeup. 


### Deliverable
You should [send us](mailto:engineering@serifhealth.com) a link to a public repository or zip file that contains at miminum:
1. The script or code used to parse the file and produce output. 
2. The setup or packaging file(s) required to bootstrap and execute your solution code
3. The output URL list.
4. A README file, explaining your solution, how long it took you to write, how long it took to run, and the tradeoffs you made along the way. 

## Expectations
### Time vs Quality
We are a small engineering team with limited resources, and often have to make hard tradeoffs to meet deadlines and make rapid forward progress. We do not want this takehome to take more than a few hours out of your day. So, please timebox coding your solution to two hours max, and know that you have the opportunity to discuss the tradeoffs you made when submitting your solution. Experienced engineers should be able to complete the coding portion in about 90 minutes, perhaps less if you have prior healthcare experience. If you think this will take you dramatically more time than that, let us know before starting the takehome so we can discuss why. 

If you finish early, we'd recommend adding additional notes or commentary to the README (e.g. discussion of performance characteristics, how you would ideally test/deploy/run your code in a production environment, feature iterations that might come next, so on), but please don't exceed the timebox doing so. 

### Language Choice
You can choose any language you want, but your solution should be portable enough to run on someone else's machine. 

### Dependencies
You can and *probably should* use dependencies (JSON parsers, type validators, etc) and libraries from public package managers in your language of choice. Again, your solution should be portable enough to run on someone else's machine, so if you leverage packaged dependencies this please make sure relevant setup instructions to install the dependencies and execute the solution are included.
