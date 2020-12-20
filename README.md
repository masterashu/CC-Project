# CC Project Report

Ram Nad 			S20180010145

Vitthal Inani 		S20180010017

Ashutosh Chauhan 	S20180010017 

 [Presentation](https://docs.google.com/presentation/d/1J36EW_ApEM9CS1lHT3hqEHCuE8XipgpgZAN-OcMfjY8/edit?usp=sharing), [Code Repository](https://github.com/masterashu/CC-Project)

## Problem Statement

Analysis of containers architecture and design, and to find something that can be improved.

## Introduction 

Over the past two years, we have seen significant changes take place that affected, and continue to affect how Containers are adopted. As we enter the new decade, we want to recap the changes and developments that we saw and offer our view of where we believe Containers are heading to in 2021.We looked into building to develop a container from scratch in a well known language Go, a low level programming language with a good concurrency for sub routing and OS level calls. We succeeded in creating local containers and following are our findings.

## Design/Methodology

A Container is a term used to define an isolated process. The process can have various types of isolation, such as storage, network, IPC, etc. We used 4 Namespaces for providing isolation:

1. Mount Namespace to run it in an isolated storage.
2. PID Namespace to hide host processes inside the container.
3. UTS Namespace to hide hostname and other details.
4. User Namespace to allow the container to be run without root privileges on the host machine.

## Results

We wrote a go program to implement our container which uses docker to download any image from docker hub and run it inside our container. The container will be isolated from other processes in terms of storage, processes, and users/groups. ![img](https://lh5.googleusercontent.com/cvS8n0ZOt44AF1uZDTMmOyN7WWocjAVb0R8x01GGLM4tjT2d34kWBCvDjY1ERn6HY5M3pVXIRFct5uhMcGqiZnWSA-o_Dd0yspfh8tKDVW6ipztY2vUuHep_ayMBjf9tjPy_7M8)![img](https://lh6.googleusercontent.com/osusWjax2w_heVqmxzCKw31EYpvxNlZZp-LAVFetrBjCw23Xxok6U07xBnQq3gq_hcDyJbMnDOE10ZxzB5WOeYDE2zG6pmX9_Ay0uelrKXjci5sX1Jxvicxi-PYpXtJ8dEX5Oxs)
